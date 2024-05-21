package imagemagick

const DefaultMorphFrames = 9

type morphOptions struct {
	frames int
}

func defaultMorphOptions() *morphOptions {
	return &morphOptions{
		frames: DefaultMorphFrames,
	}
}

// WithNumMorphFrames sets the number of frames to generate in the morph sequence. Equivalent to the -morph flag in ImageMagick.
func WithNumMorphFrames(frames int) MorphOption {
	return func(o *morphOptions) {
		o.frames = frames
	}
}

// MorphOption is a functional option for configuring an ImageMagick instance.
type MorphOption func(*morphOptions)
