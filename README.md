Usage of ./build/timelapse-linux-amd64:

Usage: ./build/timelapse-linux-amd64 [OPTIONS]

    Branch: main
    Hash: 58bc0a57ffbdbacb52c7a401cbec3b1866bb0659
    Dirty: true

A non-interactive command line utility to create smooth timlapse videos from
a set of images. This tool requires pre-installed command line tools to be
installed in order to run successfully.

Required Tools:

    ffmpeg          https://ffmpeg.org/
    imagemagick     https://imagemagick.org/index.php

Example:

    ./build/timelapse-linux-amd64 --source-directory=/path/to/image/directory

Options:

  -output-file string
        Name of the generated video file. (default "output.mp4")
  -source-dir string
        Directory containing images to be processed. These images MUST be named in lexical order. 