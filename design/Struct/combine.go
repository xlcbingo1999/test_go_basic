package Struct

import "fmt"

// 组合模式: 把一个数据结构构建成树状模式
// 树的节点需要实现同一个接口, 这样就可以快速插入和删除逻辑, 不用写更多的代码
// 树的节点可以分属不同的类别, 这样就可以区分根节点和叶子节点

type Component interface {
	Print() // 把整个子树打印出来, DFS打印
}

type Folder struct {
	components []Component
	name       int
}

func (f *Folder) Print() {
	fmt.Printf("%d ", f.name)
	for _, c := range f.components {
		c.Print() // 递归使用同一个接口!
	}
	fmt.Println()
}

type File struct {
	name int
}

func (f *File) Print() {
	fmt.Printf("%d ", f.name)
}

func RunCombine() {
	file1 := &File{name: 1}
	file2 := &File{name: 2}
	file3 := &File{name: 3}

	fol1 := &Folder{components: []Component{
		&Folder{components: []Component{file3}, name: 5}, file1, file2,
	}, name: 4}

	fol1.Print()
}
