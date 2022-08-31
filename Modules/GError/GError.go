package gerror

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

//Erorr struct...
type GError struct {
	app       fyne.App
	w         fyne.Window
	ErrorText string
}

//Создание новой ошибки
func NewGError(a fyne.App, ErrorText string) *GError {
	e := GError{a, a.NewWindow("Error"), ErrorText}
	return &e
}

//Отображение окна ошибки
func (GEr GError) ShowError() {
	GEr.w.Resize(fyne.NewSize(70, 60))
	//Перекрашивает цвет в красный
	er_lab := canvas.NewText(GEr.ErrorText, color.NRGBA{255, 0, 0, 255})

	GEr.w.SetContent(container.NewCenter(er_lab))
	GEr.w.Show()
}
