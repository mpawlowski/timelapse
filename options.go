package timelapse

import (
	"github.com/mpawlowski/timelapse/src/pkg/ffmpeg"
	"github.com/mpawlowski/timelapse/src/pkg/imagemagick"
)

type timelapseOptions struct {
	ffmpegOptions      []ffmpeg.Option
	imagemagickOptions []imagemagick.Option
}

// Option is an interface for functional options.
type Option func(*timelapseOptions)

func defaultTimelapseOptions() *timelapseOptions {
	return &timelapseOptions{
		ffmpegOptions:      []ffmpeg.Option{},
		imagemagickOptions: []imagemagick.Option{},
	}
}

// WithFFMpegVideoOptions is a functional option for setting ffmpeg options.
func WithFFMpegVideoOptions(opts ...ffmpeg.Option) Option {
	return func(t *timelapseOptions) {
		t.ffmpegOptions = append(t.ffmpegOptions, opts...)
	}
}

// WithImageMagickMorphOptions is a functional option for setting imagemagick options.
func WithImageMagickMorphOptions(opts ...imagemagick.Option) Option {
	return func(t *timelapseOptions) {
		t.imagemagickOptions = append(t.imagemagickOptions, opts...)
	}
}
