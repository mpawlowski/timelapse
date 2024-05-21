package ffmpeg

const DefaultFFMpegBinary = "ffmpeg"
const DefaultVideoBitRate = "9M"
const DefaultVideoFrameRate = 24
const DefaultVideoPixelFormat = "yuv420p"
const DefaultVideoSize = "1280x720"
const DefaultVideoCodec = "libx264"

type options struct {
	bitRate      string
	ffmPegBinary string
	frameRate    int
	pixelFormat  string
	videoSize    string
	videoCodec   string
}

func defaultFFMpegOptions() *options {
	return &options{
		bitRate:      DefaultVideoBitRate,
		ffmPegBinary: DefaultFFMpegBinary,
		frameRate:    DefaultVideoFrameRate,
		pixelFormat:  DefaultVideoPixelFormat,
		videoSize:    DefaultVideoSize,
		videoCodec:   DefaultVideoCodec,
	}
}

// WithBitRate sets the bit rate of the video. Equivalent to the -b:v flag in FFMpeg.
func WithBitRate(bitRate string) Option {
	return func(o *options) {
		o.bitRate = bitRate
	}
}

// WithCustomBinary sets the path to the FFMpeg binary to use.
func WithCustomBinary(binary string) Option {
	return func(o *options) {
		o.ffmPegBinary = binary
	}
}

// WithFrameRate sets the frame rate of the video. Equivalent to the -framerate flag in FFMpeg.
func WithFrameRate(frameRate int) Option {
	return func(o *options) {
		o.frameRate = frameRate
	}
}

// WithPixelFormat sets the pixel format of the video. Equivalent to the -pix_fmt flag in FFMpeg.
func WithPixelFormat(pixelFormat string) Option {
	return func(o *options) {
		o.pixelFormat = pixelFormat
	}
}

// WithVideoSize sets the size of the video. Equivalent to the -s:v flag in FFMpeg.
func WithVideoSize(videoSize string) Option {
	return func(o *options) {
		o.videoSize = videoSize
	}
}

// WithVideoCodec sets the codec of the video. Equivalent to the -vcodec flag in FFMpeg.
func WithVideoCodec(videoCodec string) Option {
	return func(o *options) {
		o.videoCodec = videoCodec
	}
}

// Option is a functional option for configuring an FFMpeg instance.
type Option func(*options)
