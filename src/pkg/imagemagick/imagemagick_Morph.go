package imagemagick

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/mpawlowski/timelapse/src/pkg/osutil"
)

func (i *imageMagick) Morph(ctx context.Context, from string, to string, destDir string, options ...Option) error {

	opts := defaultMorphOptions()
	for _, opt := range options {
		opt(opts)
	}

	if !osutil.CheckIfProgramExists(opts.imageMagickBinary) {
		return fmt.Errorf("ImageMagick binary not found at: %s", opts.imageMagickBinary)
	}

	dest := fmt.Sprintf("%s/%%d.jpg", destDir)

	args := []string{}
	args = append(args, from)
	args = append(args, to)
	args = append(args, "-morph", fmt.Sprintf("%d", opts.frames))
	args = append(args, dest)

	fmt.Println("exec", "convert", args)

	exec := exec.CommandContext(ctx, "convert", args...)
	exec.Stdout = os.Stdout
	exec.Stderr = os.Stderr
	return exec.Run()
}
