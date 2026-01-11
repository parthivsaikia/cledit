package cledit

type Filter struct {
	Inputs []*Filter
	Kind   string
	Args   []string
}
