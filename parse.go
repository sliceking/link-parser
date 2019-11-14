package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var xn []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		xn = append(xn, linkNodes(c)...)
	}

	return xn
}

func buildLink(n *html.Node) Link {
	var link = Link{}
	for _, a := range n.Attr {
		if a.Key == "href" {
			link.Href = a.Val
			break
		}
	}
	link.Text = text(n)

	return link
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	// if it is an element node, keep drilling
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	// whitespace is removed by using Fields then joined
	return strings.Join(strings.Fields(ret), " ")
}
