//go:build windows

package win

const (
	// InitCommonControlsEx flags
	ICC_LISTVIEW_CLASSES   = 0x00000001
	ICC_TREEVIEW_CLASSES   = 0x00000002
	ICC_BAR_CLASSES        = 0x00000004
	ICC_TAB_CLASSES        = 0x00000008
	ICC_UPDOWN_CLASS       = 0x00000010
	ICC_PROGRESS_CLASS     = 0x00000020
	ICC_HOTKEY_CLASS       = 0x00000040
	ICC_ANIMATE_CLASS      = 0x00000080
	ICC_WIN95_CLASSES      = 0x000000FF
	ICC_DATE_CLASSES       = 0x00000100
	ICC_USEREX_CLASSES     = 0x00000200
	ICC_COOL_CLASSES       = 0x00000400
	ICC_INTERNET_CLASSES   = 0x00000800
	ICC_PAGESCROLLER_CLASS = 0x00001000
	ICC_NATIVEFNTCTL_CLASS = 0x00002000
	ICC_STANDARD_CLASSES   = 0x00004000
	ICC_LINK_CLASS         = 0x00008000
)

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-initcommoncontrolsex
type INITCOMMONCONTROLSEX struct {
	Size uint32
	ICC  uint32
}

//sys InitCommonControlsEx(icc *INITCOMMONCONTROLSEX) (ok bool) = comctl32.InitCommonControlsEx