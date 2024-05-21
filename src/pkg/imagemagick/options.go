package imagemagick

const DefaultImageMagickBinary = "convert"
const DefaultMorphFrames = 9

type options struct {
	frames            int
	imageMagickBinary string
}

func defaultMorphOptions() *options {
	return &options{
		frames:            DefaultMorphFrames,
		imageMagickBinary: DefaultImageMagickBinary,
	}
}

// WithNumMorphFrames sets the number of frames to generate in the morph sequence. Equivalent to the -morph flag in ImageMagick.
func WithNumMorphFrames(frames int) Option {
	return func(o *options) {
		o.frames = frames
	}
}

// WithCustomBinary sets the path to the ImageMagick binary to use.
func WithCustomBinary(binary string) Option {
	return func(o *options) {
		o.imageMagickBinary = binary
	}
}

// Option is a functional option for configuring an ImageMagick instance.
type Option func(*options)
