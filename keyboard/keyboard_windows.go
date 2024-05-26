//go:build windows

package keyboard

import (
	"unsafe"
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

func (k *Keyboard) Start() {
	k.run()
}

func (k *Keyboard) run() {
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

	var msg MSG
	for getMessage(&msg, 0, 0, 0) != 0 {
	}

	unhookWindowsHookEx(keyboardHook)
	keyboardHook = 0
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
