// Package cli provides a command line interface for generating a timelapse video from a series of images.
package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mpawlowski/timelapse"
	"github.com/mpawlowski/timelapse/src/pkg/ffmpeg"
	"github.com/mpawlowski/timelapse/src/pkg/imagemagick"
	"github.com/mpawlowski/timelapse/src/pkg/osutil"
	"go.uber.org/fx"
)

type flags struct {
	OutputFile      string
	SourceDirectory string

	FFMpegBinary      string
	FFMpegBitRate     string
	FFMpegFrameRate   int
	FFMpegPixelFormat string
	FFMpegVideoCodec  string
	FFMpegVideoSize   string

	ImageMagickBinary         string
	ImageMagickNumMorphFrames int
}

var options flags

func parseFlags() {

	flag.StringVar(&options.OutputFile, "output-file", "output.mp4", "Name of the generated video file.")
	flag.StringVar(&options.SourceDirectory, "source-dir", "", "Directory containing images to be processed.")

	flag.StringVar(&options.FFMpegBinary, "ffmpeg-binary", ffmpeg.DefaultFFMpegBinary, "Path to the FFMpeg binary.")
	flag.StringVar(&options.FFMpegBitRate, "ffmpeg-bit-rate", ffmpeg.DefaultVideoBitRate, "Bit rate of the generated video.")
	flag.IntVar(&options.FFMpegFrameRate, "ffmpeg-frame-rate", ffmpeg.DefaultVideoFrameRate, "Frame rate of the generated video.")
	flag.StringVar(&options.FFMpegPixelFormat, "ffmpeg-pixel-format", ffmpeg.DefaultVideoPixelFormat, "Pixel format of the generated video.")
	flag.StringVar(&options.FFMpegVideoCodec, "ffmpeg-video-codec", ffmpeg.DefaultVideoCodec, "Codec used to encode the generated video.")
	flag.StringVar(&options.FFMpegVideoSize, "ffmpeg-video-size", ffmpeg.DefaultVideoSize, "Dimensions of the generated video.")

	flag.StringVar(&options.ImageMagickBinary, "imagemagick-binary", imagemagick.DefaultImageMagickBinary, "Path to the ImageMagick binary.")
	flag.IntVar(&options.ImageMagickNumMorphFrames, "imagemagick-num-morph-frames", imagemagick.DefaultMorphFrames, "Number of frames to generate between each image.")

	// Override the default help message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", HelpText())
		flag.PrintDefaults()
	}

	flag.Parse()

	// check the source directory
	if options.SourceDirectory == "" {
		fmt.Fprintf(flag.CommandLine.Output(), "Error: --source-directory is empty\n")
		flag.Usage()
		os.Exit(1)
	}

	// check and see if ffmpeg binary exists
	if !osutil.CheckIfProgramExists(options.FFMpegBinary) {
		fmt.Fprintf(flag.CommandLine.Output(), "Error: ffmpeg binary does not exist at %s\n", options.FFMpegBinary)
		flag.Usage()
		os.Exit(127)
	}

	// check and see if imagemagick binary exists
	if !osutil.CheckIfProgramExists(options.ImageMagickBinary) {
		fmt.Fprintf(flag.CommandLine.Output(), "Error: imagemagick binary does not exist at %s\n", options.ImageMagickBinary)
		flag.Usage()
		os.Exit(127)
	}
}

func buildOptions() []timelapse.Option {
	tOpts := []timelapse.Option{}

	tOpts = append(tOpts, timelapse.WithFFMpegVideoOptions(
		ffmpeg.WithCustomBinary(options.FFMpegBinary),
		ffmpeg.WithBitRate(options.FFMpegBitRate),
		ffmpeg.WithFrameRate(options.FFMpegFrameRate),
		ffmpeg.WithVideoSize(options.FFMpegVideoSize),
		ffmpeg.WithVideoCodec(options.FFMpegVideoCodec),
	))

	tOpts = append(tOpts, timelapse.WithImageMagickMorphOptions(
		imagemagick.WithCustomBinary(options.ImageMagickBinary),
		imagemagick.WithNumMorphFrames(options.ImageMagickNumMorphFrames),
	))

	return tOpts
}

func startTimelapseGeneration(lc fx.Lifecycle, sd fx.Shutdowner, t timelapse.TimelapseGenerator) {

	ctx := context.Background()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func(myCtx context.Context) {
				err := t.GenerateTimelapse(myCtx, options.SourceDirectory, options.OutputFile, buildOptions()...)
				if err != nil {
					fmt.Println("GenerateTimelapse: ", err)
					sd.Shutdown(fx.ExitCode(1))
					return
				}
				sd.Shutdown(fx.ExitCode(0))
			}(ctx)
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}

func fxOptions() []fx.Option {
	opts := []fx.Option{}
	opts = append(opts, timelapse.FxOptions()...)
	opts = append(opts, fx.Invoke(startTimelapseGeneration))
	return opts
}

func main() {
	parseFlags()
	app := fx.New(fxOptions()...)

	// start
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		panic(err)
	}

	// wait for app to finish
	signal := <-app.Wait()

	// stop
	stopCtx, stopCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer stopCancel()
	err = app.Stop(stopCtx)
	if err != nil {
		panic(err)
	}

	// exit with signal
	os.Exit(signal.ExitCode)
}
