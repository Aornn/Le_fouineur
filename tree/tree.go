package tree

import (
	"fmt"
	l "rise_perso/levenshtein"
)

type leaf struct {
	Key          rune
	CompleteWord string
	IsEnded      bool
	Leaf         Tree
}

//Tree : is the type of the prefix tree
type Tree map[rune]*leaf

//NewTree Create New tree
func NewTree() Tree {
	tree := make(Tree, 26)
	return tree
}

func recursiveSearch(query string, currRune rune, node Tree, tab []string) []string {
	for i := currRune; i <= 'z'; i++ {
		if node[i] != nil {
			if node[i].IsEnded == true && node[i].CompleteWord[0] == query[0] {
				// fmt.Println(node[i].CompleteWord + "|")
				if l.DamereauLevenshtein(query, node[i].CompleteWord) < 2 {
					// fmt.Println(node[i].CompleteWord + "|")
					tab = append(tab, node[i].CompleteWord)
					// fmt.Println(tab)
				}
			}
			tab = recursiveSearch(query, 'a', node[i].Leaf, tab)
		}
	}
	return tab
}

//SearchWord : Search for the given Word
func (tree Tree) SearchWord(query string) []string {
	node := tree
	// lenQ := len(query)
	var res []string

	currRune := rune(query[0])
	// save := node[currRune]
	out := recursiveSearch(query, currRune, node, res)
	fmt.Println(out)
	return out
}

// AddWord : adding word in tree
func (tree Tree) AddWord(query string) {
	node := tree
	for i := 0; i < len(query); i++ {
		currRune := rune(query[i])
		if i == len(query)-1 {
			if node[currRune] == nil {
				node[currRune] = &leaf{Key: currRune, CompleteWord: query, IsEnded: true, Leaf: make(Tree, 26)}
			} else {
				node[currRune].CompleteWord = query
				node[currRune].IsEnded = true
			}
		} else {
			if node[currRune] == nil {
				node[currRune] = &leaf{Key: currRune, CompleteWord: "", IsEnded: false, Leaf: make(Tree, 26)}
			}
		}
		node = node[currRune].Leaf
	}
}
