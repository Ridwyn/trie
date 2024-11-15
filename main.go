package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Children [256]*Node
	Char     string
	word     string
	graphviz string
}

type Trie struct {
	RootNode *Node
}

// Trie constructor
func NewTrie() *Trie {
	root := &Node{Char: ""}
	return &Trie{RootNode: root}
}

func (t *Trie) insert(word string) {

	current := t.RootNode
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	for i := 0; i < len(strippedWord); i++ {

		index := int(strippedWord[i])
		if index >= 256 {
			fmt.Printf("Error: Non Ascii characters not supported %v", string(strippedWord[i]))
			return
		}

		//check if current has node present at index
		if current.Children[index] == nil {
			node := &Node{Char: string(strippedWord[i])}
			if len(word)-1 == i {
				node.word = word
			}

			current.Children[index] = node

		}
		current = current.Children[index]

	}

}

func (t *Trie) ContainsWord(word string) bool {
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	current := t.RootNode
	for i := 0; i < len(strippedWord); i++ {
		index := int(strippedWord[i])
		if index >= 256 {
			fmt.Printf("Error: Non Ascii characters not supported %v", string(strippedWord[i]))
			return false
		}
		//when we reach a null path then thats last node or
		//  the decimal index of letter is nil then the word is not indexed
		if current == nil || current.Children[index] == nil {
			return false
		}

		current = current.Children[index]
	}
	return true
}

func (t *Trie) Autocomplete(word string) string {

	current := t.RootNode
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	str := word

	for i := 0; i < len(strippedWord); i++ {
		index := int(strippedWord[i])
		if index >= 256 {
			fmt.Printf("Error: Non Ascii characters not supported %v", string(strippedWord[i]))
		}
		
		if current.Children[index] == nil {
			fmt.Printf("No word for Prefix  %v", string(strippedWord[i]))
			return str
		}
		current = current.Children[index]

	}

	if len(current.word) != 0 {
		str = word
		return str
	}

	//Traverse through all node to find a word
	stack := make([]*Node, 0)
	stack = append(stack, current)
out:
	for {
		if len(stack) == 0 {
			fmt.Printf("No words for Prefix")
			break out
		}
		//get first element of stack
		current = stack[0]
		//pop off first element off stack
		stack = stack[1:]

		for i := 0; i < 256; i++ {
			if current.Children[i] != nil {
				stack = append(stack, current.Children[i])

				if len(current.Children[i].word) != 0 {
					str = current.Children[i].word
					break out
				}

			}
		}

	}

	return str

}

func (t *Trie) Print() string {
	graph := "digraph{"
	str := ""
	nodes := ""

	current := t.RootNode
	//Traverse through all nodes and children
	stack := make([]*Node, 0)
	stack = append(stack, current)

	
	count := 0
	//create graphiz node for Root
	node := fmt.Sprintf("%s%d", "node_", count)
	label := " [label=\"" + current.Char + "\"]"
	current.graphviz = node
	nodes += node + label + ";\n"

out:
	for {
		if len(stack) == 0 {
			fmt.Printf("End of Trie")
			break out
		}
		//get and pop off first element off stack
		current = stack[0]
		stack = stack[1:]

		//get graphiz node of parent and append children
		str += current.graphviz + "-> {"

		for i := 0; i < 256; i++ {
			if current.Children[i] != nil {
				//add children to stack
				stack = append(stack, current.Children[i])
				//create graphiz node od children
				cnode := fmt.Sprintf("%s%s%d", "node_", current.graphviz, i)
				clabel := " [label=\"" + current.Children[i].Char + "\"]"
				current.Children[i].graphviz = cnode
				nodes += cnode + clabel + ";\n"
				str += cnode
				if string(str[len(str)-1]) != "{" {
					str += " "
				}

			}
			if i == 255 {
				str += "}; \n"
			}
		}

		count++
	}

	graph += nodes + "\n" + str + "}"
	return graph
}

func main() {

	t := NewTrie()
	t.insert("hello")
	t.insert("hell")
	t.insert("helli")
	t.insert("bag")
	t.insert("baggage")
	t.insert("dorm")
	t.insert("doe")

	//TODO print trie using graphiz svg and png
	os.WriteFile("out.dot", []byte(t.Print()), os.ModePerm)

	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", t.ContainsWord("bag"))
	fmt.Printf("%#v\n", t.Autocomplete("hello"))

}
