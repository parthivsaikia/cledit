package cledit

// This method will gives the list of input files that are needed to be imported for a clip to work.

func (c *Clip) getInputFiles() []string {
	files := []string{}
	visitedMap := make(map[*Filter]bool)
	scan(c.Audio, visitedMap, &files)
	scan(c.Video, visitedMap, &files)
	return files
}

// This function will traverse a Filter graph and will list the ones that will have the Kind as Input:Video or Input:Audio.
// The Args of this filters will be the file names that needs to be imported.

func scan(filter *Filter, visitedMap map[*Filter]bool, resultInputs *[]string) {
	if filter == nil {
		return
	}
	if visitedMap[filter] {
		return
	}
	visitedMap[filter] = true
	if filter.Kind == OpInputAudio || filter.Kind == OpInputVideo {
		*resultInputs = append(*resultInputs, filter.Args...)
	}
	for _, neighbourFilter := range filter.Inputs {
		scan(neighbourFilter, visitedMap, resultInputs)
	}
}
