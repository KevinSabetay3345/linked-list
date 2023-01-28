package list

import (
	"errors"
	"math"
	"sort"
	"strconv"
)

func (list *List) PushBack(value int) {
	if list.First == nil {
		list.First = &Node{Value: value}
		list.Last = list.First
		list.Size++
		return
	}

	if list.First == list.Last {
		list.Last = &Node{Value: value, Prev: list.First}
		list.First.Next = list.Last
		list.Size++
		return
	}

	list.Last.Next = &Node{Value: value, Prev: list.Last}
	list.Last = list.Last.Next
	list.Size++
}

func (list *List) PushFront(value int) {
	if list.First == nil {
		list.First = &Node{Value: value}
		list.Last = list.First
		list.Size++
		return
	}

	if list.First == list.Last {
		list.First = &Node{Value: value, Next: list.Last}
		list.Last.Prev = list.First
		list.Size++
		return
	}

	list.First.Prev = &Node{Value: value, Next: list.First}
	list.First = list.First.Prev
	list.Size++
}

func (list *List) Exists(value int) bool {
	if list.First == nil {
		return false
	}

	node := list.First
	for node.Next != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}

	return node.Value == value
}

func (list *List) Position(x int) (*Node, error) {
	if x >= list.Size {
		return nil, errors.New("position " + strconv.Itoa(x) + " does not exists")
	}

	//iterates backward or forward according to size
	half := float64(list.Size) / 2.
	if x < int(math.Ceil(half)) {
		node := list.First
		for x > 0 {
			node = node.Next
			x--
		}

		return node, nil
	}

	node := list.Last
	for x < list.Size-1 {
		node = node.Prev
		x++
	}

	return node, nil

}

func (list *List) Minimum() (int, error) {
	if list.First == nil {
		return 0, errors.New("there are no elements")
	}

	min := list.First.Value
	node := list.First
	for node.Next != nil {
		if node.Value < min {
			min = node.Value
		}
		node = node.Next
	}
	if node.Value < min {
		min = node.Value
	}

	return min, nil
}

func (list *List) Maximum() (int, error) {
	if list.First == nil {
		return 0, errors.New("there are no elements")
	}

	max := list.First.Value
	node := list.First
	for node.Next != nil {
		if node.Value > max {
			max = node.Value
		}
		node = node.Next
	}
	if node.Value > max {
		max = node.Value
	}

	return max, nil
}

func (list *List) Sort() []int {
	if list.First == nil {
		return []int{}
	}

	var result []int
	node := list.First
	for node != nil {
		result = append(result, node.Value)
		node = node.Next
	}

	sort.Ints(result)

	return result
}

func (list *List) DeletePosition(pos int) error {

	if pos >= list.Size {
		return errors.New("position " + strconv.Itoa(pos) + " does not exists")
	}

	if list.Size == 1 {
		list.First = nil
		list.Last = nil
		list.Size--
		return nil
	}

	if pos == 0 {
		list.First = list.First.Next
		list.First.Prev = nil
		list.Size--
		return nil
	}

	if pos == list.Size-1 {
		list.Last = list.Last.Prev
		list.Last.Next = nil
		list.Size--
		return nil
	}

	//iterates backward or forward according to size
	half := float64(list.Size) / 2.
	if pos < int(math.Ceil(half)) {
		node := list.First
		for pos > 0 {
			node = node.Next
			pos--
		}

		node.Next.Prev = node.Prev
		node.Prev.Next = node.Next
		node = nil
		list.Size--
		return nil
	}

	node := list.Last
	for pos < list.Size {
		node = node.Prev
		pos++
	}

	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	node = nil
	list.Size--
	return nil
}

func (list *List) DeleteValue(value int) error {

	if list.First == nil {
		return errors.New("value not found")
	}

	if list.Size == 1 && list.First.Value == value {
		list.First = nil
		list.Last = nil
		list.Size--
		return nil
	}

	if list.First.Value == value {
		list.First = list.First.Next
		list.First.Prev = nil
		list.Size--
		return nil
	}

	if list.Last.Value == value {
		list.Last = list.Last.Prev
		list.Last.Next = nil
		list.Size--
		return nil
	}

	node := list.First
	for node.Next != nil {
		if node.Next.Value == value {
			node.Next = node.Next.Next
			node.Next.Prev = node
			list.Size--
			return nil
		}

		node = node.Next
	}

	return errors.New("value not found")
}
