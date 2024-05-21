// Package imagemagick provides a wrapper around the ImageMagick binary.
package imagemagick

import (
	"context"
)

// ImageMagick is an interface for interacting with a pre-installed ImageMagick binary.
type ImageMagick interface {
	// Morph generates a sequence of images that morph from one image to another.
	Morph(ctx context.Context, from string, to string, destDir string, options ...Option) error
}

type imageMagick struct{}

// NewImageMagick creates a new ImageMagick instance.
func NewImageMagick() ImageMagick {
	return &imageMagick{}
}
