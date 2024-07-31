package Construct

import "fmt"

// 可以实现: 不依赖所属类的情况下, 复制一个复杂对象
// 实现拷贝接口的类型可以是各种各样的, 不只是一种类型
// 本质上就是Golang的深拷贝构造函数

type INode interface {
	clone() INode
	print(string)
}

type File struct {
	name string
}

func (f *File) clone() INode {
	return &File{
		name: f.name,
	}
}

func (f *File) print(indentation string) {
	fmt.Printf("%s -> name: %s\n", indentation, f.name)
}

type Folder struct {
	children []INode // 这里就是依赖倒置原则, 要依赖于抽象类, 而不是具体类
	name     string
}

func (f *Folder) clone() INode {
	res := &Folder{
		children: make([]INode, len(f.children)),
		name:     f.name,
	}
	for i, c := range f.children {
		res.children[i] = c.clone() // 内部也要深度拷贝
	}
	return res
}

func (f *Folder) print(indentation string) {
	fmt.Printf("%s -> name: %s\n", indentation, f.name)
	for _, c := range f.children {
		c.print(indentation + indentation)
	}
}

func RunProtocal() {
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}

	folder1 := &Folder{
		children: []INode{file1, file2},
		name:     "folder1",
	}

	cloneFolder := folder1.clone()
	cloneFolder.print("\t")
}
