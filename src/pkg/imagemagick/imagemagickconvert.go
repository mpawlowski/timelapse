// Package imagemagick provides a wrapper around the ImageMagick binary.
package imagemagick

import (
	"context"
)

// ImageMagickConvert is an interface for interacting with a pre-installed ImageMagickConvert binary.
type ImageMagickConvert interface {
	// Morph generates a sequence of images that morph from one image to another.
	Morph(ctx context.Context, from string, to string, destDir string, options ...ConvertOption) error
}

type imageMagickConvert struct{}

// NewImageMagick creates a new ImageMagick instance.
func NewImageMagick() ImageMagickConvert {
	return &imageMagickConvert{}
}
