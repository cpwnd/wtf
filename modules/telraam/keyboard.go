package telraam

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(widget.Refresh)

	//widget.SetKeyboardChar("j", widget.Next, "Select next item")
	//widget.SetKeyboardChar("k", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("o", widget.openCameraSegment, "Open story in browser")
	//widget.SetKeyboardChar("c", widget.openComments, "Open comments in browser")

	//widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	//widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openCameraSegment, "Open story in browser")
	//widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
}
