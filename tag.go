package typewriter

// +gen slice
type Tag struct {
	Name    string
	Values  []TagValue
	Negated bool
}

type TagValue struct {
	Name           string
	Parameters     map[string]string
	typeParameters []item
}
