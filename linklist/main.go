package main

import "fmt"

// example of empty interface says nothing


type LinkedList struct {
	Value interface{}
	Next *LinkedList
}

func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList {
			Value: val,
			Next: ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}

func (ll *LinkedList) String() string {
	if ll == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v->%s", ll.Value, ll.Next)
}

func main() {
	var ll *LinkedList
	fmt.Println(ll)
	ll = ll.Insert(0, 10)
	fmt.Println(ll)
	ll = ll.Insert(0, "foo")
	fmt.Println(ll)
	ll = ll.Insert(0, 4.5)
	fmt.Println(ll)
	ll = ll.Insert(0, false)
	fmt.Println(ll)

// if we insert empty struct below, it will ruin the linked list order 
//	ll = ll.Insert(2, struct{}{})
//	fmt.Println(ll)
}
