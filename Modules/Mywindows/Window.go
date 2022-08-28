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

func (win *Window) ShowWin() {
	win.w.Show()
}
