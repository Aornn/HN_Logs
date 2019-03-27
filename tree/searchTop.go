package tree

import (
	"fmt"
	"sort"
	"time"
)

func (data NbReq) rank(size int) []orderCount {
	var order []orderCount
	var ret []orderCount
	for k, v := range data {
		order = append(order, orderCount{k, v})
	}
	sort.Slice(order, func(i, j int) bool {
		return order[i].Count > order[j].Count
	})
	for i := 0; i < size; i++ {
		ret = append(ret, orderCount{order[i].Query, order[i].Count})
	}
	return ret
}

func (ret NbReq) addintab(toInsert *NbReq) {
	for k, v := range *toInsert {
		_, ok := ret[k]
		if ok {
			ret[k] = ret[k] + v
		} else {
			ret[k] = v
		}

	}
}

func (node Tree) recursiveSearchTop(time []int, start int, ret NbReq) NbReq {
	var currTime int
	if start < 6 {
		if start <= len(time)-1 {
			currTime = time[start]
			if node[currTime] != nil {
				if node[currTime].IsEnded == true {
					ret.addintab(&node[currTime].Infos)
				}
				ret = node[currTime].Leaf.recursiveSearchTop(time, start+1, ret)
			}
		} else {
			currTime = 0
			if start == 1 || start == 2 {
				currTime = 1
			}
			for cpt := currTime; cpt < hashTime[start]; cpt++ {
				if node[cpt] != nil {
					if node[cpt].IsEnded == true {
						ret.addintab(&node[cpt].Infos)
					}
					ret = node[cpt].Leaf.recursiveSearchTop(time, start+1, ret)
				}
			}
		}
	}
	return ret
}

//CountDataTop : return the most popular query on the given timestamp
func (node Tree) CountDataTop(timest []int, size int) []orderCount {
	ret := make(NbReq)
	// var
	start := time.Now()
	ret = node.recursiveSearchTop(timest, 0, ret)
	elapsed := time.Since(start)
	fmt.Printf("Search took : %s\n======\n", elapsed)
	return ret.rank(size)
}
