package cledit

func Import(filename string) *Clip {
	video := Filter{
		Kind:   OpInputVideo,
		Args:   []string{filename},
		Inputs: nil,
	}
	audio := Filter{
		Kind:   OpInputAudio,
		Args:   []string{filename},
		Inputs: nil,
	}

	return &Clip{
		Video: &video,
		Audio: &audio,
	}
}
