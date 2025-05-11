package bst

type node struct {
	left   *node
	data   int
	right  *node
	parent *node
}

type BinaryTree struct {
	root *node
}

func New() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(value int) {
	if bt.root == nil {
		bt.root = &node{data: value}

	} else {
		current := bt.root
		for current != nil {
			if value <= current.data {
				if current.left == nil {
					current.left = &node{data: value, parent: current}
					return
				}
				current = current.left
			} else {
				if current.right == nil {
					current.right = &node{data: value, parent: current}
					return
				}
				current = current.right
			}
		}
	}
}

func (bt *BinaryTree) Find(value int) (bool, *node) {
    current := bt.root
    for current != nil {
        if value == current.data {
            return true, current
        }
        if value < current.data {
            current = current.left
        } else {
            current = current.right
        }
    }
    return false, nil
}

func (bt *BinaryTree) FindMin(arg ...*node) (int, *node) {
	current := bt.root
	if len(arg) > 0 {
		current = arg[0]
		for current.left != nil {
			current = current.left
		}
		return current.data, current
	}
	for current.left != nil {
		current = current.left
	}
	return current.data, current
}

func (bt *BinaryTree) FindMax(arg ...*node) (int, *node) {
	current := bt.root
	if len(arg) > 0 {
		current = arg[0]
		for current.right != nil {
			current = current.right
		}
		return current.data, current
	}
	for current.right != nil {
		current = current.right
	}

	return current.data, current
}


func (bt *BinaryTree) Delete(value int) {
    found, deleteNode := bt.Find(value)
    if !found {
        return
    }

	replaceNode := func(parent, oldNode, newNode *node) {
        if parent == nil {
            bt.root = newNode
        } else if parent.left == oldNode {
            parent.left = newNode
        } else {
            parent.right = newNode
        }
        if newNode != nil {
            newNode.parent = parent
        }
    }


    if deleteNode.left == nil && deleteNode.right == nil {
        replaceNode(deleteNode.parent, deleteNode, nil)
        return
    }
    if deleteNode.left == nil {
        replaceNode(deleteNode.parent, deleteNode, deleteNode.right)
        return
    }
    if deleteNode.right == nil {
        replaceNode(deleteNode.parent, deleteNode, deleteNode.left)
        return
    }
    minData, minNode := bt.FindMin(deleteNode.right)
    deleteNode.data = minData
    replaceNode(minNode.parent, minNode, minNode.right)
}
