// Package timelapse provides the ability to generate a timelapse video from a
// series of sorted images.
//
// Example Usage:
//
//	package main
//
//	import (
//		"context"
//		"fmt"
//		"os"
//
//		"github.com/mpawlowski/timelapse"
//		"github.com/mpawlowski/timelapse/src/pkg/ffmpeg"
//		"github.com/mpawlowski/timelapse/src/pkg/imagemagick"
//	)
//
//	func main() {
//		im := imagemagick.NewImageMagick()
//		ff := ffmpeg.NewFFMpeg()
//		tl := timelapse.NewTimelapseGenerator(im, ff)
//
//		ctx := context.Background()
//		sourceDir := "path/to/images"
//		outputFile := "output.mp4"
//
//		err := tl.GenerateTimelapse(ctx, sourceDir, outputFile)
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//	}
package timelapse

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/mpawlowski/timelapse/src/pkg/ffmpeg"
	"github.com/mpawlowski/timelapse/src/pkg/imagemagick"
)

type TimelapseGenerator interface {
	// GenerateTimelapse creates a timelapse video from a series of images.
	GenerateTimelapse(ctx context.Context, sourceDir string, outputFile string, options ...Option) error
}

type timelapseGenerator struct {
	im imagemagick.ImageMagick
	ff ffmpeg.FFMpeg
}

func (t *timelapseGenerator) GenerateTimelapse(ctx context.Context, sourceDir string, outputFile string, options ...Option) error {

	opts := defaultTimelapseOptions()

	for _, opt := range options {
		opt(opts)
	}

	// Ensure that the source directory exists
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		return err
	}

	sourceFiles, err := os.ReadDir(sourceDir)
	if err != nil {
		return err
	}
	sort.Slice(sourceFiles, func(i, j int) bool {
		return sourceFiles[i].Name() < sourceFiles[j].Name()
	})

	// create a temporary directory to store intermediate files
	tempDir, err := os.MkdirTemp("", "timelapse")
	if err != nil {
		return fmt.Errorf("unable to generate temp directory %s: %w", tempDir, err)
	}
	defer os.RemoveAll(tempDir)

	// create a directory to store short clips of morphed frames
	clipDir := filepath.Join(tempDir, "clips")
	err = os.Mkdir(clipDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to generate clip directory %s: %w", clipDir, err)
	}

	// create a directory to store morphed frames
	frameDir := filepath.Join(tempDir, "frames")
	err = os.Mkdir(frameDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to generate frame directory %s: %w", clipDir, err)
	}

	// loop through images and create morphed frames / clips from each image pair
	for i := 0; i < len(sourceFiles)-1; i++ {
		fromFile := sourceFiles[i]
		toFile := sourceFiles[i+1]

		// create directory for frame data
		myFrameDir := filepath.Join(frameDir, strconv.Itoa(i))
		err := os.Mkdir(myFrameDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("unable to generate frame directory %s: %w", frameDir, err)
		}

		// create morphed frames
		from := filepath.Join(sourceDir, fromFile.Name())
		to := filepath.Join(sourceDir, toFile.Name())
		err = t.im.Morph(ctx, from, to, myFrameDir, opts.imagemagickOptions...)
		if err != nil {
			return fmt.Errorf("unable to convert images %s and %s: %w", from, to, err)
		}

		// rename morphed frames to zero pad them (so they are naturally ordered)
		morphedFiles, err := os.ReadDir(myFrameDir)
		if err != nil {
			return fmt.Errorf("unable to read morphed frame directory %s: %w", myFrameDir, err)
		}
		for _, morphedFile := range morphedFiles {
			newName := fmt.Sprintf("%010s", morphedFile.Name())
			err = os.Rename(filepath.Join(myFrameDir, morphedFile.Name()), filepath.Join(myFrameDir, newName))
			if err != nil {
				return fmt.Errorf("unable to rename morphed frame %s: %w", morphedFile.Name(), err)
			}
		}

		// create video clip from morphed frames
		clipFile := path.Join(clipDir, fmt.Sprintf("%d.mp4", i))
		err = t.ff.VideoFromImages(ctx, myFrameDir, clipFile, opts.ffmpegOptions...)
		if err != nil {
			return fmt.Errorf("unable to generate video clip %s: %w", clipFile, err)
		}

		// finished with frame data so it's safe to remove
		os.RemoveAll(myFrameDir)
	}

	// create a video file for ffmpeg to concat
	clipListFileName := filepath.Join(tempDir, "cliplist.txt")
	clipListFile, err := os.Create(clipListFileName)
	if err != nil {
		return fmt.Errorf("unable to create clip list file %s: %w", clipListFileName, err)
	}
	defer clipListFile.Close()
	clipList, err := os.ReadDir(clipDir)
	if err != nil {
		return fmt.Errorf("unable to read clip directory %s: %w", clipDir, err)
	}
	for _, clip := range clipList {
		absoluteClipFilePath, err := filepath.Abs(filepath.Join(clipDir, clip.Name()))
		if err != nil {
			return fmt.Errorf("unable to get absolute path for %s: %w", clip.Name(), err)
		}
		clipListFile.WriteString(fmt.Sprintf("file '%s'\n", absoluteClipFilePath))
	}
	clipListFile.Close()

	// create final timelapse by combining the clips
	err = t.ff.Concat(ctx, clipListFileName, outputFile)
	if err != nil {
		return fmt.Errorf("unable to generate final timelapse from %s: %w", clipListFileName, err)
	}

	return nil
}

// NewTimelapseGenerator creates a new TimelapseGenerator instance.
func NewTimelapseGenerator(
	im imagemagick.ImageMagick,
	ff ffmpeg.FFMpeg,
) TimelapseGenerator {
	return &timelapseGenerator{
		im: im,
		ff: ff,
	}
}
