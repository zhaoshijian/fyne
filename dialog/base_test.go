package dialog

import (
	"image/color"
	"testing"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func TestShowCustom_ApplyTheme(t *testing.T) {
	a := test.NewApp()
	a.Settings().SetTheme(theme.DarkTheme())
	w := test.NewWindow(canvas.NewRectangle(color.Transparent))
	w.Resize(fyne.NewSize(300, 200))
	d := NewCustom("Title", "OK", widget.NewLabel("Content"), w)

	d.Show()
	test.AssertImageMatches(t, "dialog-custom-dark.png", w.Canvas().Capture())

	a.Settings().SetTheme(theme.LightTheme())
	test.AssertImageMatches(t, "dialog-custom-light.png", w.Canvas().Capture())
}
