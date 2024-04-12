package mapx

import (
	"fmt"

	sync "github.com/xlcbingo1999/test_go_basic/mapx/sync"
)

type val struct {
	s string
}

func RunSyncMapX() {
	var m sync.Map
	fmt.Println("sync.Map init status:")
	m.Dump()
	val1 := &val{"val1"}
	m.Store("key1", val1)
	fmt.Println("\nafter store key1:")
	m.Dump()

	m.Load("key2")
	fmt.Println("\nafter load key2:")
	m.Dump()

	val2 := &val{"val2"}
	m.Store("key2", val2)
	fmt.Println("\nafter store key2:")
	m.Dump()

	val3 := &val{"val3"}
	m.Store("key3", val3)
	fmt.Println("\nafter store key3:")
	m.Dump()

	m.Load("key1")
	fmt.Println("\nafter load key1:")
	m.Dump()

	m.Load("key2")
	fmt.Println("\nafter load key2:")
	m.Dump()

	m.Load("key2")
	fmt.Println("\nafter load key2 2nd:")
	m.Dump()

	m.Load("key2")
	fmt.Println("\nafter load key2 3rd:")
	m.Dump()

	/*
		// 验证update key
		val2_1 := &val{"val2_1"}
		m.Store("key2", val2_1)
		fmt.Println("\nafter update key2(in read, not in dirty):")
		m.Dump()

		val4 := &val{"val4"}
		m.Store("key4", val4)
		fmt.Println("\nafter store key4:")
		m.Dump()

		val4_1 := &val{"val4_1"}
		m.Store("key4", val4_1)
		fmt.Println("\nafter update key4(not in read, in dirty):")
		m.Dump()
	*/

	m.Delete("key2")
	fmt.Println("\nafter delete key2:")
	m.Dump()

	val4 := &val{"val4"}
	m.Store("key4", val4)
	fmt.Println("\nafter store key4:")
	m.Dump()

	m.Delete("key4")
	fmt.Println("\nafter delete key4:")
	m.Dump()

	m.Delete("key1")
	fmt.Println("\nafter delete key1:")
	m.Dump()

	/*
		val5 := &val{"val5"}
		m.Store("key5", val5)
		fmt.Println("\nafter store key5:")
		m.Dump()
	*/

	m.Load("key5")
	fmt.Println("\nafter load key5:")
	m.Dump()

	m.Load("key5")
	fmt.Println("\nafter load key5 2nd:")
	m.Dump()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(curr *TreeNode, sum int, targetSum int, currPath *[]int, res *[][]int) {
	if sum == targetSum && curr.Left == nil && curr.Right == nil {
		*res = append(*res, *currPath)
		return
	}

	if curr.Left != nil {
		*currPath = append(*currPath, curr.Left.Val)
		dfs(curr.Left, sum+curr.Left.Val, targetSum, currPath, res)
		*currPath = (*currPath)[0 : len(*currPath)-1]
	}

	if curr.Right != nil {
		*currPath = append(*currPath, curr.Right.Val)
		dfs(curr.Right, sum+curr.Right.Val, targetSum, currPath, res)
		*currPath = (*currPath)[0 : len(*currPath)-1]
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	currPath := make([]int, 0)
	dfs(root, root.Val, targetSum, &currPath, &res)
	return res
}
