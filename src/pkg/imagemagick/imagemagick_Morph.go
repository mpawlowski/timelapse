package imagemagick

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func (i *imageMagick) Morph(ctx context.Context, from string, to string, destDir string, options ...MorphOption) error {

	opts := defaultMorphOptions()
	for _, opt := range options {
		opt(opts)
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
