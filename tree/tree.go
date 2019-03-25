package tree

import (
	"fmt"
	"sort"
	"time"

	l "catcher/levenshtein"
)

type ReturnData struct {
	Data arrStudent
	Time float64
}

type Student struct {
	CompleteWord string
	Login        string `json:"login"`
	DisplayName  string `json:"displayname"`
	Level        string `json:"level"`
	Url          string `json:"url"`
	PhotoUrl     string `json:"image_url"`
}

type leaf struct {
	Key          rune
	CompleteWord string
	IsEnded      bool
	Infos        []Student
	Leaf         Tree
}

type arrStudent []Student

//Tree : is the type of the prefix tree
type Tree map[rune]*leaf

//NewTree Create New tree
func NewTree() Tree {
	tree := make(Tree, 52)
	return tree
}

func (tab arrStudent) ranking(query string) {
	sort.Slice(tab, func(i, j int) bool {
		return l.DamereauLevenshtein(tab[i].CompleteWord, query) < l.DamereauLevenshtein(tab[j].CompleteWord, query)
	})
}

func recursiveSearch(query string, currRune rune, node Tree, tab arrStudent, max float64) arrStudent {

	for i := currRune; i <= 'z'; i++ {
		if node[i] != nil {
			if node[i].IsEnded == true && node[i].CompleteWord[0] == query[0] {
				if l.DamereauLevenshtein(query, node[i].CompleteWord) < max {
					for j := 0; j < len(node[i].Infos); j++ {
						tab = append(tab, node[i].Infos[j])
					}
				}
			}
			tab = recursiveSearch(query, '-', node[i].Leaf, tab, max)
		}
		if i == '-' {
			i = 96
		}
	}
	return tab
}

//SearchWord : Search for the given Word
func (tree Tree) SearchWord(query string) ReturnData {
	start := time.Now()
	node := tree
	var res arrStudent
	var toRet ReturnData
	currRune := rune(query[0])
	if len(query) < 4 {
		toRet.Data = recursiveSearch(query, currRune, node, res, 4)
	} else if len(query) < 6 {
		toRet.Data = recursiveSearch(query, currRune, node, res, 3)
	} else {
		toRet.Data = recursiveSearch(query, currRune, node, res, 2)
	}
	toRet.Data.ranking(query)
	elapsed := time.Since(start)
	toRet.Time = float64(elapsed / time.Millisecond)
	fmt.Printf("Research took : %s\n======\n", elapsed)
	return toRet
}

// AddWord : adding word in tree
func (tree Tree) AddWord(user Student, query string) {
	node := tree
	for i := 0; i < len(query); i++ {
		currRune := rune(query[i])
		if i == len(query)-1 {
			if node[currRune] == nil {
				node[currRune] = &leaf{Key: currRune, CompleteWord: query, IsEnded: true, Leaf: make(Tree, 52)}
				node[currRune].Infos = append(node[currRune].Infos, Student{CompleteWord: query, Login: user.Login, DisplayName: user.DisplayName, Level: user.Level, Url: user.Url, PhotoUrl: user.PhotoUrl})
			} else {
				node[currRune].CompleteWord = query
				node[currRune].IsEnded = true
				node[currRune].Infos = append(node[currRune].Infos, Student{CompleteWord: query, Login: user.Login, DisplayName: user.DisplayName, Level: user.Level, Url: user.Url, PhotoUrl: user.PhotoUrl})
			}
		} else {
			if node[currRune] == nil {
				node[currRune] = &leaf{Key: currRune, CompleteWord: "", IsEnded: false, Leaf: make(Tree, 52)}
			}
		}
		node = node[currRune].Leaf
	}
}
