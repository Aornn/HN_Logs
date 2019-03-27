package tree

import (
	"fmt"
	"time"
)

func (node Tree) recursiveSearch(time []int, start int, ret int) int {
	var currTime int
	if start < 6 {
		if start <= len(time)-1 {
			currTime = time[start]
			if node[currTime] != nil {
				if node[currTime].IsEnded == true {
					ret = ret + node[currTime].Total
				}
				ret = node[currTime].Leaf.recursiveSearch(time, start+1, ret)
			}
		} else {
			currTime = 0
			if start == 1 || start == 2 {
				currTime = 1
			}
			for cpt := currTime; cpt < hashTime[start]; cpt++ {
				if node[cpt] != nil {
					if node[cpt].IsEnded == true {
						ret = ret + node[cpt].Total
					}
					ret = node[cpt].Leaf.recursiveSearch(time, start+1, ret)
				}
			}
		}
	}
	return ret
}

// CountData : Return the nuber of queries filter by the given timetsamp
func (node Tree) CountData(timest []int, size int) Datares {
	temp := 0
	var toRet Datares
	start := time.Now()
	toRet.Count = node.recursiveSearch(timest, 0, temp)
	elapsed := time.Since(start)
	fmt.Printf("Search took : %s\n======\n", elapsed)
	return toRet
}
