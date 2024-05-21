package timelapse

import (
	"github.com/mpawlowski/timelapse/src/pkg/ffmpeg"
	"github.com/mpawlowski/timelapse/src/pkg/imagemagick"
	"go.uber.org/fx"
)

// FxOptions returns an array of fx.Options for use of the timelapse package with go.uber.org/fx.
//
// Example Usage:
//
//	app := fx.New(
//		timelapse.FxOptions()...,
//	)
func FxOptions() []fx.Option {
	opts := []fx.Option{}
	opts = append(opts, fx.Provide(NewTimelapseGenerator))
	opts = append(opts, fx.Provide(imagemagick.NewImageMagick))
	opts = append(opts, fx.Provide(ffmpeg.NewFFMpeg))
	return opts
}
