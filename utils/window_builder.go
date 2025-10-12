package utils

type Window struct {
	StartIndex  int
	WindowLines []string
}

// To provide defaults for window size and overlap size
type WindowOptions struct {
	WindowSize  int
	OverlapSize int
}

func WindowBuilder(lines []string, opts *WindowOptions) ([]Window, err) {
	if opts == nil {
		opts = &WindowOptions{
			WindowSize:  300,
			OverlapSize: 100,
		}
	}

	stepSize := opts.WindowSize - opts.OverlapSize

	var windows []Window

	for startIndex := 0; startIndex < len(lines); startIndex += stepSize {

		endIndex := startIndex + opts.WindowSize
		if endIndex > len(lines) {
			endIndex = len(lines)
		}

		windowLines := lines[startIndex:endIndex]

		window := Window{
			StartIndex:  startIndex,
			WindowLines: windowLines,
		}
		windows = append(windows, window)

	}

	return windows, nil

}
