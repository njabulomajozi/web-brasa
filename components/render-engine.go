package components

import (
	"fmt"
	"strings"

	"web-bra/types"

	"golang.org/x/net/html"
)

// Widgets

// Fetch web page content

func RenderPage(URL string) {
	HTMLDocument := "<div>ME</div>"

	parsedHTML, err := parseHTMLDocument(HTMLDocument)
	if err != nil {
		panic(err)
	}

	DOMTree := &types.Document{}
	DOMTree.DOM = buildDOM(nil, parsedHTML, DOMTree)

	printNode(DOMTree.DOM)
}

func parseHTMLDocument(HTMLDocument string) (*html.Node, error) {
	// Parse HTML
	parsedHTML, err := html.Parse(strings.NewReader(HTMLDocument))
	if err != nil {
		return nil, err
	}

	// Parse CSS

	return parsedHTML, nil
}

func buildDOM(parentNode *html.Node, currentNode *html.Node, document *types.Document) *types.Node {
	DOM := &types.Node{}
	// DOM.Parent = parentNode
	// DOM.Document = document

	switch currentNode.Type {
	case html.TextNode:
		DOM.Type = "html:text"
		DOM.Text = currentNode.Data
	case html.ElementNode:
		DOM.Type = currentNode.Data
	case html.DoctypeNode:
		DOM.Type = "html:doctype"
	case html.RawNode:
		DOM.Type = "html:raw"
	}

	children := retrieveChildren(currentNode)

	results := make([]*types.Node, len(children))
	resultChan := make(chan *types.Node, len(children))

	for i, child := range children {
		go func(i int, child *html.Node) {
			childDOM := createChildDOM(currentNode, child, document, DOM)

			resultChan <- childDOM
		}(i, child)
	}

	for i := 0; i < len(children); i++ {
		results[i] = <-resultChan
	}

	close(resultChan)

	DOM.Children = append(DOM.Children, results...)

	return DOM
}

func retrieveChildren(node *html.Node) []*html.Node {
	var children []*html.Node
	if node.FirstChild == nil {
		return children
	}

	child := node.FirstChild
	children = append(children, child)

	for child.NextSibling != nil {
		child = child.NextSibling
		children = append(children, child)
	}

	return children
}

func createChildDOM(currentNode *html.Node, child *html.Node, document *types.Document, DOM *types.Node) *types.Node {
	childDOM := buildDOM(currentNode, child, document)
	childDOM.Parent = DOM

	return childDOM
}

func printNode(node *types.Node) {
	fmt.Printf("%s\t", node.Type)
	// fmt.Printf("%s\t", node.Text)
	// fmt.Printf("Children\n")
	for _, child := range node.Children {
		printNode(child)
	}
	fmt.Printf("\n\n")
}

// JS Interpreter

// Paint into window
