package mywindows

import "fyne.io/fyne/v2"

type Window struct {
	a       fyne.App
	w       fyne.Window
	content *fyne.Container
}

func New_Window(a fyne.App, WinName string, content *fyne.Container) Window {
	win := a.NewWindow(WinName)
	win.SetContent(content)
	win.CenterOnScreen()
	win.Resize(fyne.NewSize(600, 500))
	return Window{a: a, w: win, content: content}
}

func (win *Window) SetWinContent(content *fyne.Container) {
	win.w.SetContent(content)
}

func (win *Window) CloseWin() {
	win.w.Hide()
}

func (win *Window) ShowWin() {
	win.w.Show()
}
