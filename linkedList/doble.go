package linkedlist

type doubleNode[T comparable] struct{
	prev *doubleNode[T]
	data T
	next *doubleNode[T]
}

type DoubleList[T comparable] struct{
	head *doubleNode[T]
	last *doubleNode[T]
	size int
}

func (*DoubleList[T]) New() *DoubleList[T]{
	return &DoubleList[T]{}
}

func (dl *DoubleList[T]) Insert(data T){
	newNode := &doubleNode[T]{data: data}
	if dl.head == nil{
		dl.head = newNode
		dl.head.next.prev = dl.head
	} else{

	}
}