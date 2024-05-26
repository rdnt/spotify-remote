//go:build linux

package keyboard

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/MarinX/keylogger"
)

type Keyboard struct {
	keyDownHandler func(int)
	keyUpHandler   func(int)
	keys           map[int]bool
	devs           map[string]bool
	wg             *sync.WaitGroup
	cancel         context.CancelFunc
	mux            sync.Mutex
}

func New(keyDown, keyUp func(int)) *Keyboard {
	k := &Keyboard{
		keyDownHandler: keyDown,
		keyUpHandler:   keyUp,
		keys:           make(map[int]bool),
		devs:           map[string]bool{},
		wg:             &sync.WaitGroup{},
	}

	return k
}

func (k *Keyboard) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	k.cancel = cancel

	k.run(ctx)
}

func (k *Keyboard) Stop() {
	k.cancel()
	k.wg.Wait()
}

func (k *Keyboard) run(ctx context.Context) {
	k.wg.Add(1)

	go func() {
		defer k.wg.Done()

		ticker := time.NewTicker(1 * time.Second)

		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return

			case <-ticker.C:
				k.mux.Lock()

				devs := keylogger.FindAllKeyboardDevices()
				for _, dev := range devs {
					if _, ok := k.devs[dev]; !ok {
						k.devs[dev] = true

						k.track(ctx, dev)
					}
				}

				k.mux.Unlock()
			}
		}
	}()
}

func (k *Keyboard) track(ctx context.Context, dev string) {
	k.wg.Add(1)

	go func() {
		defer k.wg.Done()

		defer func() {
			k.mux.Lock()
			delete(k.devs, dev)
			k.mux.Unlock()
		}()

		l, err := keylogger.New(dev)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer l.Close()

		for {
			select {
			case <-ctx.Done():
				return

			case evt, ok := <-l.Read():
				if !ok {
					return
				}

				if evt.Type != keylogger.EvKey {
					continue
				}

				code, ok := mapKey(evt.Code)
				if !ok {
					continue
				}

				if evt.Value != 0 {
					if _, ok := k.keys[code]; !ok {
						k.keys[code] = true

						if k.keyDownHandler != nil {
							go k.keyDownHandler(code)
						}
					}
				} else {
					_, ok := k.keys[code]

					if !ok {
						if k.keyDownHandler != nil {
							go k.keyDownHandler(code)
						}
					} else {
						delete(k.keys, code)
					}

					if k.keyUpHandler != nil {
						go k.keyUpHandler(code)
					}
				}
			}
		}
	}()
}

func mapKey(key uint16) (int, bool) {
	switch key {
	case 163:
		return NextTrack, true
	case 165:
		return PrevTrack, true
	case 164:
		return PlayPause, true
	default:
		return 0, false
	}
}
