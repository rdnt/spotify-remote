//go:build windows

package keyboard

import (
	"context"
	"sync"
	"unsafe"
)

type Keyboard struct {
	keyDownHandler func(int)
	keyUpHandler   func(int)
	keys           map[int]bool
	wg             *sync.WaitGroup
	cancel         context.CancelFunc
}

func New(keyDown, keyUp func(int)) *Keyboard {
	k := &Keyboard{
		keyDownHandler: keyDown,
		keyUpHandler:   keyUp,
		keys:           make(map[int]bool),
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
		keyboardHook = setWindowsHookEx(
			WH_KEYBOARD_LL,
			func(nCode int, wparam WPARAM, lparam LPARAM) LRESULT {
				if nCode == 0 {

					kbdstruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
					code := int(kbdstruct.VkCode)
					code = mapKey(code)

					if wparam == WM_KEYDOWN || wparam == WM_SYSKEYDOWN {
						if _, ok := k.keys[code]; !ok {
							k.keys[code] = true

							if k.keyDownHandler != nil {
								go k.keyDownHandler(code)
							}
						}

					} else if wparam == WM_KEYUP {
						if _, ok := k.keys[code]; ok {
							delete(k.keys, code)

							if k.keyUpHandler != nil {
								go k.keyUpHandler(code)
							}
						}

					}
				}

				return callNextHookEx(keyboardHook, nCode, wparam, lparam)
			}, 0, 0,
		)

		go func() {
			defer k.wg.Done()

			<-ctx.Done()

			unhookWindowsHookEx(keyboardHook)
			keyboardHook = 0
		}()

		var msg MSG
		for getMessage(&msg, 0, 0, 0) != 0 {
			// TODO: how to unblock for graceful shutdown?
		}

		unhookWindowsHookEx(keyboardHook)
		keyboardHook = 0
	}()
}

func mapKey(key int) int {
	switch key {
	case VK_MEDIA_NEXT_TRACK:
		return NextTrack
	case VK_MEDIA_PREV_TRACK:
		return PrevTrack
	case VK_MEDIA_PLAY_PAUSE:
		return PlayPause
	default:
		return key
	}
}
