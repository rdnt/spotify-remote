//go:build windows

package keyboard

const (
	VK_LBUTTON             int = 0x01
	VK_RBUTTON             int = 0x02
	VK_CANCEL              int = 0x03
	VK_MBUTTON             int = 0x04
	VK_XBUTTON1            int = 0x05
	VK_XBUTTON2            int = 0x06
	VK_BACK                int = 0x08
	VK_TAB                 int = 0x09
	VK_CLEAR               int = 0x0C
	VK_RETURN              int = 0x0D
	VK_SHIFT               int = 0x10
	VK_CONTROL             int = 0x11
	VK_MENU                int = 0x12
	VK_PAUSE               int = 0x13
	VK_CAPITAL             int = 0x14
	VK_KANA                int = 0x15
	VK_IME_ON              int = 0x16
	VK_JUNJA               int = 0x17
	VK_FINAL               int = 0x18
	VK_HANJA               int = 0x19
	VK_IME_OFF             int = 0x1A
	VK_ESCAPE              int = 0x1B
	VK_CONVERT             int = 0x1C
	VK_NONCONVERT          int = 0x1D
	VK_ACCEPT              int = 0x1E
	VK_MODECHANGE          int = 0x1F
	VK_SPACE               int = 0x20
	VK_PRIOR               int = 0x21
	VK_NEXT                int = 0x22
	VK_END                 int = 0x23
	VK_HOME                int = 0x24
	VK_LEFT                int = 0x25
	VK_UP                  int = 0x26
	VK_RIGHT               int = 0x27
	VK_DOWN                int = 0x28
	VK_SELECT              int = 0x29
	VK_PRINT               int = 0x2A
	VK_EXECUTE             int = 0x2B
	VK_SNAPSHOT            int = 0x2C
	VK_INSERT              int = 0x2D
	VK_DELETE              int = 0x2E
	VK_HELP                int = 0x2F
	AK_0                   int = 0x30
	AK_1                   int = 0x31
	AK_2                   int = 0x32
	AK_3                   int = 0x33
	AK_4                   int = 0x34
	AK_5                   int = 0x35
	AK_6                   int = 0x36
	AK_7                   int = 0x37
	AK_8                   int = 0x38
	AK_9                   int = 0x39
	AK_A                   int = 0x41
	AK_B                   int = 0x42
	AK_C                   int = 0x43
	AK_D                   int = 0x44
	AK_E                   int = 0x45
	AK_F                   int = 0x46
	AK_G                   int = 0x47
	AK_H                   int = 0x48
	AK_I                   int = 0x49
	AK_J                   int = 0x4A
	AK_K                   int = 0x4B
	AK_L                   int = 0x4C
	AK_M                   int = 0x4D
	AK_N                   int = 0x4E
	AK_O                   int = 0x4F
	AK_P                   int = 0x50
	AK_Q                   int = 0x51
	AK_R                   int = 0x52
	AK_S                   int = 0x53
	AK_T                   int = 0x54
	AK_U                   int = 0x55
	AK_V                   int = 0x56
	AK_W                   int = 0x57
	AK_X                   int = 0x58
	AK_Y                   int = 0x59
	AK_Z                   int = 0x5A
	VK_LWIN                int = 0x5B
	VK_RWIN                int = 0x5C
	VK_APPS                int = 0x5D
	VK_SLEEP               int = 0x5F
	VK_NUMPAD0             int = 0x60
	VK_NUMPAD1             int = 0x61
	VK_NUMPAD2             int = 0x62
	VK_NUMPAD3             int = 0x63
	VK_NUMPAD4             int = 0x64
	VK_NUMPAD5             int = 0x65
	VK_NUMPAD6             int = 0x66
	VK_NUMPAD7             int = 0x67
	VK_NUMPAD8             int = 0x68
	VK_NUMPAD9             int = 0x69
	VK_MULTIPLY            int = 0x6A
	VK_ADD                 int = 0x6B
	VK_SEPARATOR           int = 0x6C
	VK_SUBTRACT            int = 0x6D
	VK_DECIMAL             int = 0x6E
	VK_DIVIDE              int = 0x6F
	VK_F1                  int = 0x70
	VK_F2                  int = 0x71
	VK_F3                  int = 0x72
	VK_F4                  int = 0x73
	VK_F5                  int = 0x74
	VK_F6                  int = 0x75
	VK_F7                  int = 0x76
	VK_F8                  int = 0x77
	VK_F9                  int = 0x78
	VK_F10                 int = 0x79
	VK_F11                 int = 0x7A
	VK_F12                 int = 0x7B
	VK_F13                 int = 0x7C
	VK_F14                 int = 0x7D
	VK_F15                 int = 0x7E
	VK_F16                 int = 0x7F
	VK_F17                 int = 0x80
	VK_F18                 int = 0x81
	VK_F19                 int = 0x82
	VK_F20                 int = 0x83
	VK_F21                 int = 0x84
	VK_F22                 int = 0x85
	VK_F23                 int = 0x86
	VK_F24                 int = 0x87
	VK_NUMLOCK             int = 0x90
	VK_SCROLL              int = 0x91
	VK_LSHIFT              int = 0xA0
	VK_RSHIFT              int = 0xA1
	VK_LCONTROL            int = 0xA2
	VK_RCONTROL            int = 0xA3
	VK_LMENU               int = 0xA4
	VK_RMENU               int = 0xA5
	VK_BROWSER_BACK        int = 0xA6
	VK_BROWSER_FORWARD     int = 0xA7
	VK_BROWSER_REFRESH     int = 0xA8
	VK_BROWSER_STOP        int = 0xA9
	VK_BROWSER_SEARCH      int = 0xAA
	VK_BROWSER_FAVORITES   int = 0xAB
	VK_BROWSER_HOME        int = 0xAC
	VK_VOLUME_MUTE         int = 0xAD
	VK_VOLUME_DOWN         int = 0xAE
	VK_VOLUME_UP           int = 0xAF
	VK_MEDIA_NEXT_TRACK    int = 0xB0
	VK_MEDIA_PREV_TRACK    int = 0xB1
	VK_MEDIA_STOP          int = 0xB2
	VK_MEDIA_PLAY_PAUSE    int = 0xB3
	VK_LAUNCH_MAIL         int = 0xB4
	VK_LAUNCH_MEDIA_SELECT int = 0xB5
	VK_LAUNCH_APP1         int = 0xB6
	VK_LAUNCH_APP2         int = 0xB7
	VK_OEM_1               int = 0xBA
	VK_OEM_PLUS            int = 0xBB
	VK_OEM_COMMA           int = 0xBC
	VK_OEM_MINUS           int = 0xBD
	VK_OEM_PERIOD          int = 0xBE
	VK_OEM_2               int = 0xBF
	VK_OEM_3               int = 0xC0
	VK_OEM_4               int = 0xDB
	VK_OEM_5               int = 0xDC
	VK_OEM_6               int = 0xDD
	VK_OEM_7               int = 0xDE
	VK_OEM_8               int = 0xDF
	VK_OEM_102             int = 0xE2
	VK_PROCESSKEY          int = 0xE5
	VK_PACKET              int = 0xE7
	VK_ATTN                int = 0xF6
	VK_CRSEL               int = 0xF7
	VK_EXSEL               int = 0xF8
	VK_EREOF               int = 0xF9
	VK_PLAY                int = 0xFA
	VK_ZOOM                int = 0xFB
	VK_PA1                 int = 0xFD
	VK_OEM_CLEAR           int = 0xFE
)

var keyNames = map[int]string{
	VK_LBUTTON:             "Left mouse button",
	VK_RBUTTON:             "Right mouse button",
	VK_CANCEL:              "Control-break processing",
	VK_MBUTTON:             "Middle mouse button (three-button mouse)",
	VK_XBUTTON1:            "X1 mouse button",
	VK_XBUTTON2:            "X2 mouse button",
	VK_BACK:                "⟵",
	VK_TAB:                 "TAB",
	VK_CLEAR:               "CLEAR",
	VK_RETURN:              "↵",
	VK_SHIFT:               "⇪",
	VK_CONTROL:             "^",
	VK_MENU:                "ALT",
	VK_PAUSE:               "PAUSE",
	VK_CAPITAL:             "CAPS LOCK",
	VK_KANA:                "IME Kana mode",
	VK_IME_ON:              "IME On",
	VK_JUNJA:               "IME Junja mode",
	VK_FINAL:               "IME final mode",
	VK_HANJA:               "IME Hanja mode",
	VK_IME_OFF:             "IME Off",
	VK_ESCAPE:              "⎋",
	VK_CONVERT:             "IME convert",
	VK_NONCONVERT:          "IME nonconvert",
	VK_ACCEPT:              "IME accept",
	VK_MODECHANGE:          "IME mode change request",
	VK_SPACE:               "␣",
	VK_PRIOR:               "PAGE UP",
	VK_NEXT:                "PAGE DOWN",
	VK_END:                 "END",
	VK_HOME:                "HOME",
	VK_LEFT:                "LEFT ARROW",
	VK_UP:                  "UP ARROW",
	VK_RIGHT:               "RIGHT ARROW",
	VK_DOWN:                "DOWN ARROW",
	VK_SELECT:              "SELECT",
	VK_PRINT:               "PRINT",
	VK_EXECUTE:             "EXECUTE",
	VK_SNAPSHOT:            "PRINT SCREEN",
	VK_INSERT:              "INS",
	VK_DELETE:              "DEL",
	VK_HELP:                "HELP",
	AK_0:                   "0",
	AK_1:                   "1",
	AK_2:                   "2",
	AK_3:                   "3",
	AK_4:                   "4",
	AK_5:                   "5",
	AK_6:                   "6",
	AK_7:                   "7",
	AK_8:                   "8",
	AK_9:                   "9",
	AK_A:                   "A",
	AK_B:                   "B",
	AK_C:                   "C",
	AK_D:                   "D",
	AK_E:                   "E",
	AK_F:                   "F",
	AK_G:                   "G",
	AK_H:                   "H",
	AK_I:                   "I",
	AK_J:                   "J",
	AK_K:                   "K",
	AK_L:                   "L",
	AK_M:                   "M",
	AK_N:                   "N",
	AK_O:                   "O",
	AK_P:                   "P",
	AK_Q:                   "Q",
	AK_R:                   "R",
	AK_S:                   "S",
	AK_T:                   "T",
	AK_U:                   "U",
	AK_V:                   "V",
	AK_W:                   "W",
	AK_X:                   "X",
	AK_Y:                   "Y",
	AK_Z:                   "Z",
	VK_LWIN:                "Left Windows key (Natural keyboard)",
	VK_RWIN:                "Right Windows key (Natural keyboard)",
	VK_APPS:                "Applications key (Natural keyboard)",
	VK_SLEEP:               "Computer Sleep",
	VK_NUMPAD0:             "Numeric keypad 0",
	VK_NUMPAD1:             "Numeric keypad 1",
	VK_NUMPAD2:             "Numeric keypad 2",
	VK_NUMPAD3:             "Numeric keypad 3",
	VK_NUMPAD4:             "Numeric keypad 4",
	VK_NUMPAD5:             "Numeric keypad 5",
	VK_NUMPAD6:             "Numeric keypad 6",
	VK_NUMPAD7:             "Numeric keypad 7",
	VK_NUMPAD8:             "Numeric keypad 8",
	VK_NUMPAD9:             "Numeric keypad 9",
	VK_MULTIPLY:            "Multiply",
	VK_ADD:                 "Add",
	VK_SEPARATOR:           "Separator",
	VK_SUBTRACT:            "Subtract",
	VK_DECIMAL:             "Decimal",
	VK_DIVIDE:              "Divide",
	VK_F1:                  "F1",
	VK_F2:                  "F2",
	VK_F3:                  "F3",
	VK_F4:                  "F4",
	VK_F5:                  "F5",
	VK_F6:                  "F6",
	VK_F7:                  "F7",
	VK_F8:                  "F8",
	VK_F9:                  "F9",
	VK_F10:                 "F10",
	VK_F11:                 "F11",
	VK_F12:                 "F12",
	VK_F13:                 "F13",
	VK_F14:                 "F14",
	VK_F15:                 "F15",
	VK_F16:                 "F16",
	VK_F17:                 "F17",
	VK_F18:                 "F18",
	VK_F19:                 "F19",
	VK_F20:                 "F20",
	VK_F21:                 "F21",
	VK_F22:                 "F22",
	VK_F23:                 "F23",
	VK_F24:                 "F24",
	VK_NUMLOCK:             "NUM LOCK",
	VK_SCROLL:              "SCROLL LOCK",
	VK_LSHIFT:              "⇪",
	VK_RSHIFT:              "⇪",
	VK_LCONTROL:            "^",
	VK_RCONTROL:            "^",
	VK_LMENU:               "Left MENU",
	VK_RMENU:               "Right MENU",
	VK_BROWSER_BACK:        "Browser Back",
	VK_BROWSER_FORWARD:     "Browser Forward",
	VK_BROWSER_REFRESH:     "Browser Refresh",
	VK_BROWSER_STOP:        "Browser Stop",
	VK_BROWSER_SEARCH:      "Browser Search",
	VK_BROWSER_FAVORITES:   "Browser Favorites",
	VK_BROWSER_HOME:        "Browser Start and Home",
	VK_VOLUME_MUTE:         "Volume Mute",
	VK_VOLUME_DOWN:         "Volume Down",
	VK_VOLUME_UP:           "Volume Up",
	VK_MEDIA_NEXT_TRACK:    "Next Track",
	VK_MEDIA_PREV_TRACK:    "Previous Track",
	VK_MEDIA_STOP:          "Stop Media",
	VK_MEDIA_PLAY_PAUSE:    "Play/Pause Media",
	VK_LAUNCH_MAIL:         "Start Mail",
	VK_LAUNCH_MEDIA_SELECT: "Select Media",
	VK_LAUNCH_APP1:         "Start Application 1",
	VK_LAUNCH_APP2:         "Start Application 2",
	VK_OEM_1:               ":",
	VK_OEM_PLUS:            "For any country/region, the '+'",
	VK_OEM_COMMA:           "For any country/region, the ','",
	VK_OEM_MINUS:           "For any country/region, the '-'",
	VK_OEM_PERIOD:          "For any country/region, the '.'",
	VK_OEM_2:               "Used for miscellaneous characters; it can vary by keyboard. For the US standard keyboard, the '/?'",
	VK_OEM_3:               "Used for miscellaneous characters; it can vary by keyboard. For the US standard keyboard, the '`~'",
	VK_OEM_4:               "Used for miscellaneous characters; it can vary by keyboard. For the US standard keyboard, the '[{'",
	VK_OEM_5:               "Used for miscellaneous characters; it can vary by keyboard. For the US standard keyboard, the '|'",
	VK_OEM_6:               "Used for miscellaneous characters; it can vary by keyboard. For the US standard keyboard, the ']}'",
	VK_OEM_7:               "Used for miscellaneous characters; it can vary by keyboard. For the US standard keyboard, the 'single-quote/double-quote'",
	VK_OEM_8:               "Used for miscellaneous characters; it can vary by keyboard.",
	VK_OEM_102:             "The <> keys on the US standard keyboard, or the | key on the non-US 102-key keyboard",
	VK_PROCESSKEY:          "IME PROCESS",
	VK_PACKET:              "Used to pass Unicode characters as if they were keystrokes. The VK_PACKET key is the low word of a 32-bit Virtual Key value used for non-keyboard input methods. For more information, see Remark in KEYBDINPUT, SendInput, WM_KEYDOWN, and WM_KEYUP",
	VK_ATTN:                "Attn",
	VK_CRSEL:               "CrSel",
	VK_EXSEL:               "ExSel",
	VK_EREOF:               "Erase EOF",
	VK_PLAY:                "Play",
	VK_ZOOM:                "Zoom",
	VK_PA1:                 "PA1",
	VK_OEM_CLEAR:           "Clear",
}

func KeyName(key int) string {
	if name, ok := keyNames[key]; ok {
		return name
	}

	return "Unknown"
}
