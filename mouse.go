package desktop

import "github.com/fyne-io/fyne"
import "github.com/fyne-io/fyne/canvas"

import wmtheme "github.com/fyne-io/desktop/theme"

var mouse *canvas.Image

func newMouse() fyne.CanvasObject {
	mouse = canvas.NewImageFromFile(wmtheme.PointerDefault.CachePath())
	mouse.Options.RepeatEvents = true

	if isEmbedded() {
		// hide the mouse cursor as the parent desktop will paint one
		mouse.Resize(fyne.NewSize(0, 0))
		return mouse
	}

	mouse.Resize(fyne.NewSize(24, 24))

	return mouse
}
