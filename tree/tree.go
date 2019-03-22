package tree

import (
	"fmt"
	"sort"

	l "catcher/levenshtein"
)

type Student struct {
	Login       string `json:"login"`
	DisplayName string `json:"displayname"`
	Level       string `json:"level"`
	Url         string `json:"url"`
	PhotoUrl    string `json:"image_url"`
}

type leaf struct {
	Key          rune
	CompleteWord string
	IsEnded      bool
	Infos        Student
	Leaf         Tree
}

type ArrStudent []Student

//Tree : is the type of the prefix tree
type Tree map[rune]*leaf

//NewTree Create New tree
func NewTree() Tree {
	tree := make(Tree, 52)
	return tree
}

func (tab ArrStudent) ranking(query string) {
	fmt.Println(tab)
	sort.Slice(tab, func(i, j int) bool {
		return l.DamereauLevenshtein(tab[i].Login, query) < l.DamereauLevenshtein(tab[j].Login, query)
	})
}

func recursiveSearch(query string, currRune rune, node Tree, tab ArrStudent, max float64) ArrStudent {
	for i := currRune; i <= 'z'; i++ {
		if node[i] != nil {
			if node[i].IsEnded == true && node[i].CompleteWord[0] == query[0] {
				if l.DamereauLevenshtein(query, node[i].CompleteWord) < max {
					tab = append(tab, Student{Login: node[i].Infos.Login, DisplayName: node[i].Infos.DisplayName, Level: node[i].Infos.Level, Url: node[i].Infos.Url, PhotoUrl: node[i].Infos.PhotoUrl})
				}
			}
			tab = recursiveSearch(query, '!', node[i].Leaf, tab, max)
		}
	}
	return tab
}

//SearchWord : Search for the given Word
func (tree Tree) SearchWord(query string) ArrStudent {
	node := tree
	var res ArrStudent
	var out ArrStudent
	currRune := rune(query[0])
	if len(query) < 5 {
		out = recursiveSearch(query, currRune, node, res, 6)
	} else {
		out = recursiveSearch(query, currRune, node, res, 4)
	}
	out.ranking(query)
	fmt.Println(out)
	return out
}

// AddWord : adding word in tree
func (tree Tree) AddWord(user Student) {
	node := tree
	query := user.Login
	for i := 0; i < len(query); i++ {
		currRune := rune(query[i])
		if i == len(query)-1 {
			if node[currRune] == nil {
				node[currRune] = &leaf{Key: currRune, CompleteWord: query, IsEnded: true, Infos: Student{Login: query, DisplayName: user.DisplayName, Level: user.Level, Url: user.Url, PhotoUrl: user.PhotoUrl}, Leaf: make(Tree, 52)}
			} else {
				node[currRune].CompleteWord = query
				node[currRune].IsEnded = true
				node[currRune].Infos = Student{Login: query, DisplayName: user.DisplayName, Level: user.Level, Url: user.Url, PhotoUrl: user.PhotoUrl}
			}
		} else {
			if node[currRune] == nil {
				node[currRune] = &leaf{Key: currRune, CompleteWord: "", IsEnded: false, Infos: Student{}, Leaf: make(Tree, 52)}
			}
		}
		node = node[currRune].Leaf
	}
}