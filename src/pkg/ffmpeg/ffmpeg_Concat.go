package ffmpeg

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/mpawlowski/timelapse/src/pkg/osutil"
)

// Concat implements FFMpeg. bash equivalent
func (f *ffmpeg) Concat(ctx context.Context, clipListFile string, outputFile string, options ...Option) error {

	opts := defaultFFMpegOptions()
	for _, opt := range options {
		opt(opts)
	}

	if !osutil.CheckIfProgramExists(opts.ffmPegBinary) {
		return fmt.Errorf("ImageMagick binary not found at: %s", opts.ffmPegBinary)
	}

	// ffmpeg -f concat -safe 0 -i clips.txt -c copy output.mp4
	args := []string{}
	args = append(args, "-f", "concat")
	args = append(args, "-safe", "0")
	args = append(args, "-i", clipListFile)
	args = append(args, "-c", "copy")
	args = append(args, "-y")
	args = append(args, outputFile)

	exec := exec.CommandContext(ctx, "ffmpeg", args...)
	exec.Stdout = os.Stdout
	exec.Stderr = os.Stdin
	return exec.Run()
}
