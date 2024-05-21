package timelapse

import (
	"github.com/mpawlowski/timelapse/src/pkg/ffmpeg"
	"github.com/mpawlowski/timelapse/src/pkg/imagemagick"
)

type timelapseOptions struct {
	ffmpegOptions      []ffmpeg.VideoOption
	imagemagickOptions []imagemagick.MorphOption
}

// Option is an interface for functional options.
type Option func(*timelapseOptions)

func defaultTimelapseOptions() *timelapseOptions {
	return &timelapseOptions{
		ffmpegOptions:      []ffmpeg.VideoOption{},
		imagemagickOptions: []imagemagick.MorphOption{},
	}
}

// WithFFMpegVideoOptions is a functional option for setting ffmpeg options.
func WithFFMpegVideoOptions(opts ...ffmpeg.VideoOption) Option {
	return func(t *timelapseOptions) {
		t.ffmpegOptions = append(t.ffmpegOptions, opts...)
	}
}

// WithImageMagickMorphOptions is a functional option for setting imagemagick options.
func WithImageMagickMorphOptions(opts ...imagemagick.MorphOption) Option {
	return func(t *timelapseOptions) {
		t.imagemagickOptions = append(t.imagemagickOptions, opts...)
	}
}
