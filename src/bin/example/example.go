package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mpawlowski/timelapse"
	"github.com/mpawlowski/timelapse/src/pkg/ffmpeg"
	"github.com/mpawlowski/timelapse/src/pkg/imagemagick"
)

func main() {
	im := imagemagick.NewImageMagick()
	ff := ffmpeg.NewFFMpeg()
	tl := timelapse.NewTimelapseGenerator(im, ff)

	ctx := context.Background()
	sourceDir := "path/to/images"
	outputFile := "output.mp4"

	err := tl.GenerateTimelapse(ctx, sourceDir, outputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
