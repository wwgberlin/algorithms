package tree

import "fmt"

type (
	Node struct {
		Left  *Node
		Right *Node
		Value int
	}
)


//----------------------------------------
//		Code to print the tree
//----------------------------------------

const (
	newLine      = "\n"
	emptySpace   = "    "
	middleItem   = "├── "
	continueItem = "│   "
	lastItem     = "└── "
	nilItem      = "nil"
)

type (
	printer struct {
	}

	// Printer is printer interface
	Printer interface {
		Print(*Node) string
	}
)

// Text return text of root name
func Text(n *Node) string {
	if n != nil {
		return fmt.Sprintf("%d", n.Value)
	}
	return nilItem
}

// Items return slice of tree nodes
func Items(n *Node) []*Node {
	if n != nil {
		return []*Node{n.Left, n.Right}
	}
	return nil
}

// Print return string of tree
func PrintTree(n *Node) string {
	if n != nil {
		return newPrinter().Print(n)
	}
	return nilItem
}

func newPrinter() Printer {
	return &printer{}
}

func (p *printer) Print(n *Node) string {
	return Text(n) + newLine + p.printItems(Items(n), []bool{})
}

func (p *printer) printText(text string, spaces []bool, last bool) string {
	var result string
	for _, space := range spaces {
		if space {
			result += emptySpace
		} else {
			result += continueItem
		}
	}

	indicator := middleItem
	if last {
		indicator = lastItem
	}

	return result + indicator + text + newLine
}

func (p *printer) printItems(t []*Node, spaces []bool) string {
	var result string
	for i, f := range t {
		last := i == len(t)-1
		result += p.printText(Text(f), spaces, last)
		if len(Items(f)) > 0 {
			spacesChild := append(spaces, last)
			result += p.printItems(Items(f), spacesChild)
		}
	}
	return result
}
