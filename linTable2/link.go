package main

import (
	"fmt"
)

type LinkNode struct{
	data interface{}
	next *LinkNode
}

type Link  struct{
	head *LinkNode
	tail *LinkNode
}

func (p *Link) AddNodeBegin(data interface{}){
	node := &LinkNode{
		data : data,
		next : nil,
	}

	if (p.head == nil && p.tail == nil){
		p.head = node
		p.tail = node
		return
	}

	node.next = p.head
	p.head = node
}

func (p *Link) AddNodeEnd(data interface{}){
	node := &LinkNode{
		data : data,
		next : nil,
	}

	if (p.head == nil && p.tail == nil){
		p.head = node
		p.tail = node
		return
	}

	p.tail.next = node
	p.tail = node
}

func (p *Link) DisplayLink(){
	it := p.head
	for it != nil {
		fmt.Println(it.data)
		it = it.next
	}
}

