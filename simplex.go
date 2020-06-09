// package simplex (simple X) provides a simplified interface
// to package xgbutil
package simplex

import (
	"image"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xwindow"
)

type WindowWrapper struct {
	XImage *xgraphics.Image
	Window *xwindow.Window
	X      *xgbutil.XUtil
}

func (ww WindowWrapper) Redraw() {
	ww.XImage.XDraw()
	ww.XImage.XPaint(ww.Window.Id)
}

func (ww WindowWrapper) RedrawFromImage(img image.Image) {
	xgraphics.Blend(ww.XImage, img, image.Point{})
	ww.Redraw()

}

func NewWindow(WindowTitle string, WindowContents image.Image) WindowWrapper {
	X, err := xgbutil.NewConn()
	if err != nil {
		panic(err)
	}

	ximg := xgraphics.NewConvert(X, WindowContents)

	window := ximg.XShowExtra(WindowTitle, true)
	return WindowWrapper{ximg, window, X}
	//wid, _ := xproto.NewWindowId(X)
	//screen := xproto.Setup(X).DefaultScreen(X)

	//xproto.MapWindow(X, wid)
	/*for {
		ev, xerr := X.WaitForEvent()
		if ev == nil && xerr == nil {
			fmt.Println("Both event and error are nil. Exiting...")
			return
		}

		if ev != nil {
			fmt.Printf("Event: %s\n", ev)
		}
		if xerr != nil {
			fmt.Printf("Error: %s\n", xerr)
		}
	}*/
}
