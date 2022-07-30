package linked

import (
	"errors"
	"fmt"
)

// 头节点
type Head struct {
	Next   *Node
	Length uint32
}

// 节点
type Node struct {
	Data interface{}
	Next *Node
}

// new 一个单向链表 返回头节点
func New(req []interface{}) *Head {

	head := &Head{Next: nil, Length: 0}
	if len(req) == 0 {
		return head
	}
	var lastNode *Node
	for i, v := range req {
		node := &Node{
			Data: v,
		}
		if i == 0 {
			head.Next = node
			lastNode = node
		} else {
			lastNode.Next = node
			lastNode = node
		}
		head.Length++
	}
	return head
}

func (l *Head) Len() uint32 {
	if l == nil {
		return 0
	}
	return l.Length
}

func (l *Head) newNode(data interface{}) *Node {
	// TODO 是否需要判断数据是否一样
	return &Node{
		Data: data,
	}
}

// 选择位置插入节点
func (l *Head) Insert(index uint32, data interface{}) error {
	if l == nil {
		return errors.New("请先初始化链表")
	}

	if l.Length < index {
		return errors.New("请选择合理位置插入")
	}
	newNode := l.newNode(data)

	if index == 0 { // 头节点插入
		newNode.Next = l.Next
		l.Next = newNode
	} else { // 非头节点插入
		var lastNode *Node = l.Next
		for i := 1; i < int(index); i++ {
			lastNode = lastNode.Next
		}

		newNode.Next = lastNode.Next
		lastNode.Next = newNode
	}
	l.Length++
	return nil
}

// 开始位置插入节点
func (l *Head) InsertToStart(data interface{}) error {
	return l.Insert(0, data)
}

// 结束位置插入节点
func (l *Head) InsertToEnd(data interface{}) error {
	return l.Insert(l.Length, data)
}

// 选择位置删除节点
func (l *Head) Delete(index uint32) error {
	if l == nil {
		return errors.New("请先初始化链表")
	}
	if l.Length < index {
		return errors.New("请选择合理位置删除")
	}
	if index == 0 { // 删除第一个节点
		l.Next = l.Next.Next
	} else {
		var lastNode *Node = l.Next
		for i := 1; i < int(index); i++ {
			lastNode = lastNode.Next
		}
		lastNode.Next = lastNode.Next.Next
	}
	l.Length--
	return nil
}

// // 选择节点删除节点
// func (l *Head) DeleteNode(data interface{}) error {

// 	return l.Insert(l.Length, data)
// }

// // 选择节点在链表位置
// func (l *Head) DeleteNode(data interface{}) error {

// 	return l.Insert(l.Length, data)
//}

// 顺序打印链表各个节点的值
func (l *Head) PrintfNode() {

	node := l.Next

	for i := 0; i < int(l.Len()); i++ {

		fmt.Printf("节点位置：%d; 节点值：%v;\n", i, node.Data)

		node = node.Next
	}
}
