package tree

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//Tree : is the type of the prefix tree and of the leaf
type Tree map[int]*leaf

//NbReq : is a map who count the numbers of time a request has been done
type NbReq map[string]int

var hashTime = map[int]int{
	1: 13,
	2: 32,
	3: 24,
	4: 60,
	5: 60,
}

type leaf struct {
	Key          int
	CompleteWord []int
	IsEnded      bool
	Infos        NbReq
	Total        int
	Leaf         Tree
}

type orderCount struct {
	Query string
	Count int
}

type Datares struct {
	Count int
}

//NewTree Create New tree
func NewTree() Tree {
	tree := make(Tree)
	return tree
}

// AddWord : adding word in tree
func (tree Tree) AddWord(query []int, data string) {
	node := tree
	for i := 0; i < len(query); i++ {
		currRune := query[i]
		if i == len(query)-1 {
			// fmt.Println(data)
			if node[currRune] == nil {
				node[currRune] = &leaf{Key: currRune, CompleteWord: query, IsEnded: true, Infos: make(NbReq), Leaf: make(Tree)}
				node[currRune].Infos[data] = 1
				node[currRune].Total = 1
			} else {
				_, ok := node[currRune].Infos[data]
				if ok {
					node[currRune].Infos[data]++
				} else {
					node[currRune].Infos[data] = 1
				}
				node[currRune].Total++
			}
		} else {
			if node[currRune] == nil {
				node[currRune] = &leaf{Key: currRune, Leaf: make(Tree)}
			}
		}
		node = node[currRune].Leaf
	}
}

//IndexFile : Index the given file in the tree
func IndexFile(tree Tree, path string) {
	start := time.Now()
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var newQ []int
		q := scanner.Text()
		query := strings.Split(q, "\t")[0]
		data := strings.Split(q, "\t")[1]
		query = strings.ReplaceAll(query, "-", " ")
		query = strings.ReplaceAll(query, ":", " ")
		tempQuery := strings.Split(query, " ")
		for i := 0; i < len(tempQuery); i++ {
			val, _ := strconv.Atoi(tempQuery[i])
			newQ = append(newQ, val)
		}
		if len(newQ) == 6 {
			tree.AddWord(newQ, data)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Indexation took : %s\n======\n", elapsed)
	fmt.Println("Indexing Done")
	fmt.Println("listen on :8080")
}
