package ffmpeg

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
)

func (f *ffmpeg) VideoFromImages(ctx context.Context, sourceDir string, outputFile string, options ...VideoOption) error {

	opts := defaultFFMpegOptions()
	for _, opt := range options {
		opt(opts)
	}

	args := []string{}
	args = append(args, "-framerate", fmt.Sprintf("%d", opts.frameRate))
	args = append(args, "-pattern_type", "glob")
	args = append(args, "-i", path.Join(sourceDir, "*.jpg"))
	args = append(args, "-s:v", opts.videoSize)
	args = append(args, "-b:v", opts.bitRate)
	args = append(args, "-vcodec", opts.videoCodec)
	args = append(args, "-pix_fmt", opts.pixelFormat)
	args = append(args, outputFile)

	fmt.Println("exec", "ffmpeg", args)

	exec := exec.CommandContext(ctx, "ffmpeg", args...)
	exec.Stdout = os.Stdout
	exec.Stderr = os.Stderr
	return exec.Run()

}
