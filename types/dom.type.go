package types

type Window struct {
}

type Document struct {
	// URL   string
	// Title string
	// Head  *Node
	// Body  *Node
	DOM *Node
}

type Node struct {
	Parent     *Node // *html.Node
	Type       string
	Attributes map[string]string
	Text       string
	Children   []*Node
}
