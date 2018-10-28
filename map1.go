package main

import (
	"fmt"
	"sort"
	"os"
	"bufio"
)

func main(){
	ages := make(map[string]int)
	ages2 := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	ages["jack"] = 30
	
	ages2["james"] = 34
	ages2["owen"] = 26
	ages2["spider"] = 36

	var names []string

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	
	if equal(ages, ages2) {
		fmt.Println("two maps have the same key and value")
	}else {
		fmt.Println("two maps have different key and value")
	}
}

//判断两个map是否相等
func equal(src, dst map[string]int) bool{
	if len(src) != len(dst) {
		return false
	}

	for key, srcValue := range src {
		if dstValue, ok := dst[key]; !ok || srcValue != dstValue {
			return false
		}
	}
	return true
}

//通过map实现set
func initSet(){
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
}

//通过map实现graph
func InitGraph(){
	var graph = make(map[string]map[string]bool)
	addEdge(graph, "beijing", "shanghai")
	if hasEdge(graph, "beijing", "shanghai") {
		fmt.Println("OK")
	}
}

func addEdge(graph map[string]map[string]bool, from, to string){
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(graph map[string]map[string]bool, from, to string) bool {
	return graph[from][to]
}





