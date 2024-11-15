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
}

type Trie struct {
	RootNode *Node
}

// Trie constructor
func NewTrie() *Trie {
	root := &Node{Char: "\000"}
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

func (t *Trie) autocomplete(word string) string {

	current := t.RootNode
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	str := word

	for i := 0; i < len(strippedWord); i++ {
		index := int(strippedWord[i])
		if index >= 256 {
			fmt.Printf("Error: Non Ascii characters not supported %v", string(strippedWord[i]))
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
		//pop off first element off stack
		current = stack[0]
		//use copy to reuse space of initial stack insteaf of = which may reassign
		copy(stack, stack[1:])

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

func main() {


	t := NewTrie()
	t.insert("hello")
	t.insert("hell")
	t.insert("helli")
	t.insert("bag")
	t.insert("doe")



	//TODO print trie using graphiz svg and png
	// os.WriteFile("out.dot", []byte(trie), os.ModePerm)

	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", t.ContainsWord("bag"))
	fmt.Printf("%#v\n", t.autocomplete("hello"))

}
