package cledit

import "fmt"

type Clip struct {
	Video *Filter
	Audio *Filter
}

func (c *Clip) Scale(width int, height int) *Clip {
	heightStr := fmt.Sprintf("%d", height)
	widthStr := fmt.Sprintf("%d", width)
	newVideo := &Filter{
		Kind:   "scale",
		Inputs: []*Filter{c.Video},
		Args:   []string{"x=" + widthStr, "y=" + heightStr},
	}
	return &Clip{
		Video: newVideo,
		Audio: c.Audio,
	}
}
