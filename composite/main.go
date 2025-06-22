package main

import "fmt"

type Component interface {
	Display(indent string)
}

// Leaf
type File struct {
	Name string
}

func (f *File) Display(indent string) {
	fmt.Println(indent + "- " + f.Name)
}

// Composite
type Folder struct {
	Name     string
	Children []Component
}

func (f *Folder) Display(indent string) {
	fmt.Println(indent + "+ " + f.Name)
	for _, child := range f.Children {
		child.Display(indent + "  ")
	}
}

func (f *Folder) Add(child Component) {
	f.Children = append(f.Children, child)
}

func main() {
	// ファイルを作成
	file1 := &File{Name: "file1.txt"}
	file2 := &File{Name: "file2.txt"}
	file3 := &File{Name: "file3.txt"}

	// サブフォルダを作成
	subFolder1 := &Folder{Name: "sub1"}
	subFolder2 := &Folder{Name: "sub2"}

	subFolder1.Add(file1)
	subFolder1.Add(file2)
	subFolder2.Add(file3)

	// ルートフォルダに追加
	root := &Folder{Name: "root"}
	root.Add(subFolder1)
	root.Add(subFolder2)

	// ツリー全体を表示
	root.Display("")
}
