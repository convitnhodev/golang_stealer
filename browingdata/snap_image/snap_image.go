package snapimage

import (
	"image"

	"github.com/kbinani/screenshot"
)

type SnapImage struct {
	Img *image.RGBA
}

func (c *SnapImage) Name() string {
	return "snap_image"
}

func (c *SnapImage) Parse(_ []byte) error {
	img, err := screenshot.CaptureDisplay(0)
	if err != nil {
		return err
	}

	c = &SnapImage{
		Img: img,
	}

	return nil

}
