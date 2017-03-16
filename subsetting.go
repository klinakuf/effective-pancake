package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, playground")
	
	#number of connections per client
	subset_size := 3	
	
	backends := []int{0,1,2,3,4,5,6,7,8,9,10,11}
	clients := []int{0,1,2,3,4,5,6,7,8,9}
	
	for c := range clients {
		result := subset(backends,c,subset_size)
		fmt.Printf("%v \n", result)
	}
	
}

func subset(backends []int, client_id int, subset_size int) (result []int) {
	var subset_count = len(backends) / subset_size
	var round = client_id / subset_count
	fmt.Printf("Round %v \n",round)
	fmt.Printf("Client %v \n",client_id)

	var r = rand.New(rand.NewSource(int64(round)))
	
	var backends_shuffled = make([]int, len(backends))
	var perm = r.Perm(len(backends))
	for i, v := range perm {
		backends_shuffled[v] = backends[i]
	}

        fmt.Printf("Shuffled Backends %v \n",backends_shuffled)

	var subset_id = client_id % subset_count
	var start = subset_id * subset_size
	return backends_shuffled[start : start+subset_size]
}






