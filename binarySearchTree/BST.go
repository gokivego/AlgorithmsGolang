/*
Implementation of a binary search tree operations (Insert, Search, Delete)
*/

package main

import ("fmt"
		//"container/list"
)

type Element struct {
	
	// The value in stored in the element
	Value interface{}
	left, right *Element
	
	// The binary search tree to which the element belongs
	bst *Bst
}

type Bst struct {
	root Element
}

// Creates a new Bst Element and initializes it
func New() *Bst {
	return new(Bst).Init()
} 

func (b *Bst) Init() *Bst {
	b.root.right = nil
	b.root.left = nil
	return b
}

func (b *Bst) Insert(val interface{}) {
	temp := &b.root
	for true {
		if (temp.Value == nil) {
			temp.Value = val
			temp.bst = b
			return
		} else if (val.(int) >= temp.Value.(int) && temp.right == nil) {
			temp.right = new(Element)
			temp.right.Value = val
			temp.right.bst = b
			return
		} else if (val.(int) < temp.Value.(int) && temp.left == nil) {
			temp.left = new(Element)
			temp.left.Value = val
			temp.left.bst = b
			return
		} else if (val.(int) >= temp.Value.(int) && temp.right != nil) {
			temp = temp.right
		} else {
			temp = temp.left
		}
	}	
}

// Returns if a value exists in the Binary Search Tree
func (b *Bst) Search(val interface{}) bool {
	temp := &b.root
	for true {
		if (temp.Value.(int) == val.(int)) {
			return true
		} else if (val.(int) >= temp.Value.(int) && temp.right != nil) {
			temp = temp.right
		} else if (val.(int) < temp.Value.(int) && temp.left != nil) {
			temp = temp.left
		} else {
			return false
		}
	}
	return false
}

func (b *Bst) Remove(val interface{}) {
	
}

// Inorder Traversal (Left, Root, Right)
func (b *Bst) Inorder(root *Element) {
	if (root == nil) {
		return
	}
	b.Inorder(root.left)
	fmt.Println(root.Value)
	b.Inorder(root.right)
}

func (b *Bst) Preorder(root *Element) {
	if (root == nil) {
		return
	}
	fmt.Println(root.Value)
	b.Preorder(root.left)
	b.Preorder(root.right)
}

func (b *Bst) Postorder(root *Element) {
	if (root == nil) {
		return
	}
	b.Postorder(root.left)
	b.Postorder(root.right)
	fmt.Println(root.Value)
}

func (b *Bst) Levelorder() {

}

/*
func ArrayToBST(array []int) {

}
*/

func main() {
	// Instantiate a random BST with the root element
	array := []int{5,10,12,7,6,21,35,45,22,100}
	b := New()
	for _, val := range array {
		b.Insert(val)
	}
	fmt.Println(b.Search(7))
	b.Inorder(&b.root)
	b.Preorder(&b.root)
	b.Postorder(&b.root)
	//fmt.Println(b.root)
	//fmt.Println(b.root.right)
}
