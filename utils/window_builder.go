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

func WindowBuilder(lines []string, opts *WindowOptions) []Window {
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

		// Create a copy of the slice to avoid referencing the original slice
		// which may lead to unexpected behavior if the original slice changes
		// after the window is created.
		windowLines := make([]string, endIndex-startIndex)

		// Copy the relevant lines into the new slice
		// The copy() function fills the first slice (destination) with values from the second slice (source).
		copy(windowLines, lines[startIndex:endIndex])

		window := Window{
			StartIndex:  startIndex,
			WindowLines: windowLines,
		}
		windows = append(windows, window)

	}

	return windows

}
