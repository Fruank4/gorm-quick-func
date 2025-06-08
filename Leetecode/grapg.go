package Leetecode

var graph [][]byte
var visited [][]bool

func numIslands(grid [][]byte) int {
	graph = grid
	visited = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}
	nums := 0
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] == '1' && !visited[i][j] {
				nums++
				dfsVisit(i, j)
			}
		}
	}
	return nums
}

func dfsVisit(i, j int) {
	if i < 0 || i >= len(graph) || j < 0 || j >= len(graph[i]) {
		return
	}
	if visited[i][j] || graph[i][j] == '0' {
		return
	}
	visited[i][j] = true
	dfsVisit(i-1, j)
	dfsVisit(i+1, j)
	dfsVisit(i, j-1)
	dfsVisit(i, j+1)
}

/*
*
	[[2,1,1],
	 [1,1,0],
	 [0,1,1]]
// 3是一分钟，4是2分钟，5是3分钟
*/

func orangesRotting(grid [][]int) int {

	queue := make([][]int, 0)
	offer := func(node []int) {
		queue = append(queue, node)
	}

	poll := func() []int {
		node := queue[0]
		queue = queue[1:]
		return node
	}

	var freshCount = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				offer([]int{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	var minutes = 0
	for len(queue) > 0 && freshCount > 0 {
		minutes++
		count := len(queue)
		for i := 0; i < count; i++ {
			node := poll()
			x := node[0]
			y := node[1]

			A := roting(x-1, y, grid)
			if A != nil {
				offer(A)
				freshCount--
			}
			B := roting(x+1, y, grid)
			if B != nil {
				offer(B)
				freshCount--
			}
			C := roting(x, y-1, grid)
			if C != nil {
				offer(C)
				freshCount--
			}
			D := roting(x, y+1, grid)
			if D != nil {
				offer(D)
				freshCount--
			}
		}

	}
	if freshCount > 0 {
		return -1
	}

	return minutes
}

func roting(i, j int, grid [][]int) []int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return nil
	}

	if grid[i][j] == 1 {
		grid[i][j] = 2
		return []int{i, j}
	}
	return nil
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	inDegree := make([]int, numCourses)
	edges := make([][]int, numCourses)
	for _, edge := range prerequisites {
		edges[edge[1]] = append(edges[edge[1]], edge[0])
		inDegree[edge[0]]++
	}

	queue := make([]int, 0)
	offer := func(node int) {
		queue = append(queue, node)
	}
	poll := func() int {
		node := queue[0]
		queue = queue[1:]
		return node
	}

	for k, v := range inDegree {
		if v == 0 {
			offer(k)
		}
	}

	var canDo = 0
	for len(queue) != 0 {
		node := poll()
		canDo++
		for _, val := range edges[node] {
			inDegree[val]--
			if inDegree[val] == 0 {
				offer(val)
			}
		}

	}

	return canDo == numCourses
}

type Trie struct {
	mp   map[string]bool
	tree *TrieNode
}

type TrieNode struct {
	child []*TrieNode
}

func ConstructorTrie() Trie {
	return Trie{
		mp: make(map[string]bool),
		tree: &TrieNode{
			child: make([]*TrieNode, 26),
		},
	}

}

func (this *Trie) Insert(word string) {
	this.mp[word] = true
	insert(word, this.tree)

}

func (this *Trie) StartsWith(prefix string) bool {
	return startWith(prefix, this.tree)
}

func (this *Trie) Search(word string) bool {
	return this.mp[word]
}

func insert(word string, node *TrieNode) {
	if len(word) == 0 {
		return
	}
	childNode := node.child[word[0]-'a']
	if childNode == nil {
		node.child[word[0]-'a'] = &TrieNode{
			child: make([]*TrieNode, 26),
		}
	}

	insert(word[1:], node.child[word[0]-'a'])
}

func startWith(prefix string, node *TrieNode) bool {
	if node == nil {
		return false
	}
	if len(prefix) == 0 {
		return true
	}
	if node.child[prefix[0]-'a'] == nil {
		return false
	}

	return startWith(prefix[1:], node.child[prefix[0]-'a'])
}
