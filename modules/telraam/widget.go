package telraam

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.KeyboardWidget
	view.ScrollableWidget

	cameras  []Camera
	settings *Settings
	err      error
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	var cameras []Camera
	for _, id := range utils.ToInts(widget.settings.cameraIds) {
		camera, e := GetCamera(id)
		if e == nil {
			cameras = append(cameras, camera)
		} else {
			widget.err = e
		}
	}

	widget.cameras = cameras
	widget.SetItemCount(len(cameras))

	widget.Render()
}

// Render sets up the widget data for redrawing to the screen
func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s - cameras", widget.CommonSettings().Title)

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if len(widget.cameras) == 0 {
		return title, "No cameras to display", false
	}

	var str string
	for idx, camera := range widget.cameras {
		//u, _ := url.Parse(camera.segment_id)

		row := fmt.Sprintf(
			`[%s]%2d. %s [lightblue](%s)[white]`,
			widget.RowColor(idx),
			idx+1,
			camera.Mac,
			strings.TrimPrefix(camera.SegmentLink(), "www."),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(camera.Mac))
	}

	return title, str, false
}

func (widget *Widget) openCameraSegment() {
	camera := widget.selectedCamera()
	if camera != nil {
		utils.OpenFile(camera.SegmentLink())
	}
}

func (widget *Widget) selectedCamera() *Camera {
	var camera *Camera

	sel := widget.GetSelected()
	if sel >= 0 && widget.cameras != nil && sel < len(widget.cameras) {
		camera = &widget.cameras[sel]
	}

	return camera
}
