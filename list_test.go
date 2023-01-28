package list

import (
	"strconv"
	"testing"
)

func TestEmptyList(t *testing.T) {
	var l List

	checkNoElems(l, t)

	sorted := l.Sort()
	if len(sorted) > 0 {
		t.Fatalf("Expected %d, Got %d", 0, len(sorted))
	}
	value, err := l.Maximum()
	if err == nil {
		t.Errorf("Expected %s, Got %v", "there are no elements", nil)
	}
	if value != 0 {
		t.Errorf("Expected %d, Got %d", 0, value)
	}
	value, err = l.Minimum()
	if err == nil {
		t.Errorf("Expected %s, Got %v", "there are no elements", nil)
	}
	if value != 0 {
		t.Errorf("Expected %d, Got %d", 0, value)
	}

	for i := 0; i < 10; i++ {
		err := l.DeletePosition(i)
		if err == nil {
			t.Errorf("Expected %s, Got %v", "position "+strconv.Itoa(i)+" does not exists", err)
		}
		err = l.DeleteValue(i)
		if err == nil {
			t.Errorf("Expected %s, Got %v", "value not found", err)
		}

		if l.Exists(i) {
			t.Errorf("Expected %t, Got %t", false, l.Exists(i))
		}

		node, err := l.Position(i)
		if err == nil {
			t.Errorf("Expected %s, Got %s", "position "+strconv.Itoa(i)+" does not exists", err)
		}
		if node != nil {
			t.Errorf("Expected %v, Got %+v", nil, node)
		}
	}
}

func TestOneElemList(t *testing.T) {
	//Insert
	var l List
	l.PushBack(4)
	checkOneElem(l, 4, t)

	var l2 List
	l2.PushFront(4)
	checkOneElem(l2, 4, t)

	//Delete
	err := l.DeletePosition(1)
	if err == nil {
		t.Errorf("Expected %s, Got %v", "position 1 does not exists", err)
	}
	err = l.DeletePosition(0)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkNoElems(l, t)

	l.PushFront(4)
	checkOneElem(l, 4, t)

	err = l.DeleteValue(3)
	if err == nil {
		t.Errorf("Expected %s, Got %v", "value not found", err)
	}
	err = l.DeleteValue(4)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkNoElems(l, t)

	l.PushBack(4)
	checkOneElem(l, 4, t)

	//get
	if !l.Exists(4) {
		t.Errorf("Expected %t, Got %t", true, false)
	}
	if l.Exists(2) {
		t.Errorf("Expected %t, Got %t", false, true)
	}

	node, err := l.Position(1)
	if err == nil {
		t.Errorf("Expected %s, Got %s", "position 1 does not exists", err)
	}
	if node != nil {
		t.Errorf("Expected %v, Got %+v", nil, node)
	}
	node, err = l.Position(0)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if node != l.First {
		t.Errorf("Expected %+v, Got %v", l.First, node)
	}

	sorted := l.Sort()
	if len(sorted) != 1 {
		t.Fatalf("Expected %d, Got %d", 1, len(sorted))
	}
	if sorted[0] != 4 {
		t.Errorf("Expected %d, Got %d", 4, sorted[0])
	}

	x, err := l.Maximum()
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if x != 4 {
		t.Errorf("Expected %d, Got %d", 4, x)
	}
	x, err = l.Minimum()
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if x != 4 {
		t.Errorf("Expected %d, Got %d", 4, x)
	}

	checkOneElem(l, 4, t)

}

func TestTwoElemsList(t *testing.T) {
	//Insert
	var l List
	l.PushFront(4)
	l.PushBack(5)
	checkTwoElems(l, 4, 5, t)

	//Delete
	err := l.DeletePosition(2)
	if err == nil {
		t.Errorf("Expected %s, Got %v", "position 2 does not exists", err)
	}
	err = l.DeletePosition(1)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkOneElem(l, 4, t)

	l.PushFront(3)
	err = l.DeletePosition(1)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkOneElem(l, 3, t)

	l.PushBack(6)
	checkTwoElems(l, 3, 6, t)

	err = l.DeletePosition(0)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkOneElem(l, 6, t)

	err = l.DeletePosition(0)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkNoElems(l, t)

	//get
	l.PushBack(7)
	l.PushBack(4)

	if !l.Exists(4) {
		t.Errorf("Expected %t, Got %t", true, !l.Exists(4))
	}
	if l.Exists(2) {
		t.Errorf("Expected %t, Got %t", false, l.Exists(2))
	}

	node, err := l.Position(2)
	if err == nil {
		t.Errorf("Expected %s, Got %s", "position 2 does not exists", err)
	}
	if node != nil {
		t.Errorf("Expected %v, Got %+v", nil, node)
	}

	node, err = l.Position(1)
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if node != l.Last {
		t.Errorf("Expected %+v, Got %+v", l.Last, node)
	}
	if node.Value != 4 {
		t.Errorf("Expected %d, Got %d", 4, node.Value)
	}

	node, err = l.Position(0)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if node != l.First {
		t.Errorf("Expected %+v, Got %v", l.First, node)
	}
	if node.Value != 7 {
		t.Errorf("Expected %d, Got %d", 7, node.Value)
	}

	sorted := l.Sort()
	if len(sorted) != 2 {
		t.Fatalf("Expected %d, Got %d", 2, len(sorted))
	}
	if sorted[1] != 7 {
		t.Errorf("Expected %d, Got %d", 7, sorted[1])
	}
	if sorted[0] != 4 {
		t.Errorf("Expected %d, Got %d", 4, sorted[0])
	}

	x, err := l.Maximum()
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if x != 7 {
		t.Errorf("Expected %d, Got %d", 7, x)
	}

	x, err = l.Minimum()
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if x != 4 {
		t.Errorf("Expected %d, Got %d", 4, x)
	}

	checkTwoElems(l, 7, 4, t)

}

func TestManyElemsList(t *testing.T) {
	// Insert
	var threeElemsList List
	threeElemsList.PushBack(8)
	threeElemsList.PushBack(2)
	threeElemsList.PushBack(5)
	checkManyElems(threeElemsList, 3, 8, 5, t)

	if threeElemsList.First.Next != threeElemsList.Last.Prev {
		t.Errorf("Expected %s, Got %s", "first.next = last.prev", "first.next != last.prev")
	}

	var l List
	l.PushBack(8)
	l.PushBack(3)
	l.PushBack(5)
	l.PushBack(1)
	checkManyElems(l, 4, 8, 1, t)

	// Delete
	err := l.DeletePosition(0)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkManyElems(l, 3, 3, 1, t)

	err = l.DeleteValue(3)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkTwoElems(l, 5, 1, t)

	l.PushBack(15)
	checkManyElems(l, 3, 5, 15, t)

	l.DeletePosition(2)
	checkTwoElems(l, 5, 1, t)

	l.DeletePosition(1)
	checkOneElem(l, 5, t)

	l.DeletePosition(0)
	checkNoElems(l, t)

	// Get
	l.PushBack(4)
	l.PushBack(15)
	l.PushBack(1)

	if !l.Exists(15) {
		t.Errorf("Expected %t, Got %t", true, false)
	}
	if l.Exists(2) {
		t.Errorf("Expected %t, Got %t", false, true)
	}

	node, err := l.Position(0)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if node != l.First {
		t.Errorf("Expected %+v, Got %v", l.First, node)
	}
	if node.Value != 4 {
		t.Errorf("Expected %d, Got %d", 4, node.Value)
	}

	sorted := l.Sort()
	if len(sorted) != 3 {
		t.Fatalf("Expected %d, Got %d", 3, len(sorted))
	}
	if sorted[0] != 1 {
		t.Errorf("Expected %d, Got %d", 1, sorted[0])
	}
	if sorted[1] != 4 {
		t.Errorf("Expected %d, Got %d", 4, sorted[1])
	}
	if sorted[2] != 15 {
		t.Errorf("Expected %d, Got %d", 15, sorted[2])
	}

	x, err := l.Maximum()
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if x != 15 {
		t.Errorf("Expected %d, Got %d", 15, x)
	}

	x, err = l.Minimum()
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if x != 1 {
		t.Errorf("Expected %d, Got %d", 1, x)
	}

	checkManyElems(l, 3, 4, 1, t)
}

// size must be bigger than 10
var size int = 1000

func TestStress(t *testing.T) {
	var l List
	var node *Node
	var err error

	for i := 0; i < size; i++ {
		//reverse order list: i <-> i-1 <-> ... <-> 1 <-> 0
		l.PushFront(i)

		if !l.Exists(i) {
			t.Errorf("Expected %t, Got %t", true, false)
		}
		if l.Exists(i + 1) {
			t.Errorf("Expected %t, Got %t", false, true)
		}

		node, err = l.Position(0)
		if err != nil {
			t.Errorf("Expected %v, Got %s", nil, err.Error())
		}
		if node != l.First {
			t.Errorf("Expected %+v, Got %v", l.First, node)
		}
		if node.Value != i {
			t.Errorf("Expected %d, Got %d", i, node.Value)
		}
		node, err = l.Position(i)
		if err != nil {
			t.Errorf("Expected %v, Got %s", nil, err.Error())
		}
		if node != l.Last {
			t.Errorf("Expected %+v, Got %v", l.Last, node)
		}
		if node.Value != 0 {
			t.Errorf("Expected %d, Got %d", 0, node.Value)
		}

		sorted := l.Sort()
		if len(sorted) != i+1 {
			t.Fatalf("Expected %d, Got %d", i+1, len(sorted))
		}
		for j := 0; j <= i; j++ {
			if sorted[j] != j {
				t.Errorf("Expected %d, Got %d", j, sorted[j])
			}
		}

		x, err := l.Maximum()
		if err != nil {
			t.Errorf("Expected %v, Got %s", nil, err.Error())
		}
		if x != i {
			t.Errorf("Expected %d, Got %d", i, x)
		}
		x, err = l.Minimum()
		if err != nil {
			t.Errorf("Expected %v, Got %s", nil, err.Error())
		}
		if x != 0 {
			t.Errorf("Expected %d, Got %d", 0, x)
		}
		checkManyElems(l, i+1, i, 0, t)
	}

	err = l.DeletePosition(size/2 + 3)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = l.DeletePosition(size/2 - 3)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	err = l.DeleteValue(size/2 + 1)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = l.DeleteValue(size/2 - 1)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	err = l.DeleteValue(0)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	checkManyElems(l, size-5, size-1, 1, t)

	var l2 List
	for i := 0; i < size; i++ {
		l2.PushBack(i)
		if !l2.Exists(i) {
			t.Errorf("Expected %t, Got %t", true, false)
		}
		if l2.Exists(i + 1) {
			t.Errorf("Expected %t, Got %t", false, true)
		}

		x, err := l2.Maximum()
		if err != nil {
			t.Errorf("Expected %v, Got %s", nil, err.Error())
		}
		if x != i {
			t.Errorf("Expected %d, Got %d", i, x)
		}
		x, err = l2.Minimum()
		if err != nil {
			t.Errorf("Expected %v, Got %s", nil, err.Error())
		}
		if x != 0 {
			t.Errorf("Expected %d, Got %d", 0, x)
		}
		checkManyElems(l2, i+1, 0, i, t)
	}

	node, err = l2.Position(size/2 - 3)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if node.Value != size/2-3 {
		t.Errorf("Expected %d, Got %d", size/2-3, node.Value)
	}
	node, err = l2.Position(size/2 + 3)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if node.Value != size/2+3 {
		t.Errorf("Expected %d, Got %d", size/2+3, node.Value)
	}

}

func checkNoElems(l List, t *testing.T) {
	if l.Size != 0 {
		t.Errorf("Expected %d, Got %d", 0, l.Size)
	}
	if l.First != nil {
		t.Errorf("Expected %v, Got %+v", nil, l.First)
	}
	if l.Last != nil {
		t.Errorf("Expected %v, Got %+v", nil, l.Last)
	}
	err := l.DeleteValue(0)
	if err == nil {
		t.Errorf("Expected %s, Got %v", "value not found", err)
	}
}

func checkOneElem(l List, value int, t *testing.T) {
	if l.Size != 1 {
		t.Errorf("Expected %d, Got %d", 1, l.Size)
	}
	if l.First != l.Last {
		t.Errorf("Expected %s, Got %s", "first = last", "first != last")
	}
	if l.First == nil {
		t.Fatalf("Expected %s, Got %s", "first != nil", "first = nil")
	}
	if l.First.Value != value {
		t.Errorf("Expected %d, Got %d", value, l.First.Value)
	}
}

func checkTwoElems(l List, v1, v2 int, t *testing.T) {
	if l.Size != 2 {
		t.Errorf("Expected %d, Got %d", 2, l.Size)
	}
	if l.First == nil {
		t.Fatalf("Expected %s, Got %s", "first != nil", "first = nil")
	}
	if l.Last == nil {
		t.Fatalf("Expected %s, Got %s", "last != nil", "last = nil")
	}
	if l.First == l.Last {
		t.Errorf("Expected %s, Got %s", "first != last", "first = last")
	}
	if l.First.Value != v1 {
		t.Errorf("Expected %d, Got %d", v1, l.First.Value)
	}
	if l.Last.Value != v2 {
		t.Errorf("Expected %d, Got %d", v2, l.Last.Value)
	}
	if l.First.Next != l.Last {
		t.Errorf("Expected %s, Got %s", "first.next = last", "first.next != last")
	}
	if l.First != l.Last.Prev {
		t.Errorf("Expected %s, Got %s", "first = last.prev", "first != last.prev")
	}
}

func checkManyElems(l List, cantNodos, firstVal, lastVal int, t *testing.T) {
	if l.Size != cantNodos {
		t.Errorf("Expected %d, Got %d", cantNodos, l.Size)
	}
	if l.First == nil {
		t.Fatalf("Expected %s, Got %s", "first != nil", "first = nil")
	}
	if l.Last == nil {
		t.Fatalf("Expected %s, Got %s", "last != nil", "last = nil")
	}
	if l.First.Value != firstVal {
		t.Errorf("Expected %d, Got %d", firstVal, l.First.Value)
	}
	if l.Last.Value != lastVal {
		t.Errorf("Expected %d, Got %d", lastVal, l.Last.Value)
	}
}
