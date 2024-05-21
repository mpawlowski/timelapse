package imagemagick

const DefaultImageMagickConvertBinary = "convert"
const DefaultMorphFrames = 9

type convertOptions struct {
	frames            int
	imageMagickBinary string
}

func defaultConvertOptions() *convertOptions {
	return &convertOptions{
		frames:            DefaultMorphFrames,
		imageMagickBinary: DefaultImageMagickConvertBinary,
	}
}

// WithNumMorphFrames sets the number of frames to generate in the morph sequence. Equivalent to the -morph flag in ImageMagick.
func WithNumMorphFrames(frames int) ConvertOption {
	return func(o *convertOptions) {
		o.frames = frames
	}
}

// WithCustomBinary sets the path to the ImageMagick binary to use.
func WithCustomBinary(binary string) ConvertOption {
	return func(o *convertOptions) {
		o.imageMagickBinary = binary
	}
}

// ConvertOption is a functional option for configuring an ImageMagick instance.
type ConvertOption func(*convertOptions)
