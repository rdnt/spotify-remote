//go:build !windows

package keyboard

import (
	"fmt"

	"github.com/MarinX/keylogger"
)

type Keyboard struct {
	keyDownHandler func(int)
	keyUpHandler   func(int)
	keys           map[int]bool
}

func New(keyDown, keyUp func(int)) *Keyboard {
	k := &Keyboard{
		keyDownHandler: keyDown,
		keyUpHandler:   keyUp,
		keys:           make(map[int]bool),
	}

	return k
}

func (k *Keyboard) Capture() {
	go k.capture()
}

var fs []func()

func (k *Keyboard) Stop() {
	for _, f := range fs {
		f()
	}
}

func (k *Keyboard) capture() {
	// TODO: track connects/disconnects
	devs := keylogger.FindAllKeyboardDevices()

	for _, dev := range devs {
		l, err := keylogger.New(dev)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fs = append(fs, func() { l.Close() })

		go func() {
			for evt := range l.Read() {
				if evt.Type != keylogger.EvKey {
					continue
				}

				code := mapKey(evt.Code)

				if evt.Value == 1 {
					if _, ok := k.keys[code]; !ok {
						k.keys[code] = true

						if k.keyDownHandler != nil {
							go k.keyDownHandler(code)
						}
					}
				} else {
					if _, ok := k.keys[code]; ok {
						delete(k.keys, code)

						if k.keyUpHandler != nil {
							go k.keyUpHandler(code)
						}
					}
				}
			}
		}()
	}
}

func mapKey(key uint16) int {
	switch key {
	case 163:
		return NextTrack
	case 165:
		return PrevTrack
	case 164:
		return PlayPause
	default:
		return int(key)
	}
}
