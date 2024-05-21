package ffmpeg

const DefaultVideoBitRate = "9M"
const DefaultVideoFrameRate = 24
const DefaultVideoPixelFormat = "yuv420p"
const DefaultVideoSize = "1280x720"
const DefaultVideoCodec = "libx264"

type ffmpegoptions struct {
	bitRate     string
	frameRate   int
	pixelFormat string
	videoSize   string
	videoCodec  string
}

func defaultFFMpegOptions() *ffmpegoptions {
	return &ffmpegoptions{
		bitRate:     DefaultVideoBitRate,
		frameRate:   DefaultVideoFrameRate,
		pixelFormat: DefaultVideoPixelFormat,
		videoSize:   DefaultVideoSize,
		videoCodec:  DefaultVideoCodec,
	}
}

// WithBitRate sets the bit rate of the video. Equivalent to the -b:v flag in FFMpeg.
func WithBitRate(bitRate string) VideoOption {
	return func(o *ffmpegoptions) {
		o.bitRate = bitRate
	}
}

// WithFrameRate sets the frame rate of the video. Equivalent to the -framerate flag in FFMpeg.
func WithFrameRate(frameRate int) VideoOption {
	return func(o *ffmpegoptions) {
		o.frameRate = frameRate
	}
}

// WithPixelFormat sets the pixel format of the video. Equivalent to the -pix_fmt flag in FFMpeg.
func WithPixelFormat(pixelFormat string) VideoOption {
	return func(o *ffmpegoptions) {
		o.pixelFormat = pixelFormat
	}
}

// WithVideoSize sets the size of the video. Equivalent to the -s:v flag in FFMpeg.
func WithVideoSize(videoSize string) VideoOption {
	return func(o *ffmpegoptions) {
		o.videoSize = videoSize
	}
}

// WithVideoCodec sets the codec of the video. Equivalent to the -vcodec flag in FFMpeg.
func WithVideoCodec(videoCodec string) VideoOption {
	return func(o *ffmpegoptions) {
		o.videoCodec = videoCodec
	}
}

// VideoOption is a functional option for configuring an FFMpeg instance.
type VideoOption func(*ffmpegoptions)
