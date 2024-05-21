// Package ffmpeg provides a wrapper around the FFMpeg binary.
package ffmpeg

import "context"

// FFMpeg is an interface for interacting with a pre-installed FFMpeg binary.
type FFMpeg interface {
	// VideoFromImages generates a video from a directory of images, ordered by
	// filename. The images MUST be in a format that FFMpeg can read. Ensure to
	// zero pad the filenames if you want them to be in order.
	VideoFromImages(ctx context.Context, sourceDir string, outputFile string, options ...Option) error
	// Concat concatenates all the video files in a directory into a single video
	// file. The files MUST be in a format that FFMpeg can read, MUST be named in
	// lexical order, and MUST have the same codec and resolution.
	Concat(ctx context.Context, sourceDir string, outputFile string, options ...Option) error
}

type ffmpeg struct{}

// NewFFMpeg creates a new FFMpeg instance.
func NewFFMpeg() FFMpeg {
	return &ffmpeg{}
}
