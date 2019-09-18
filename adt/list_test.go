package adt

import (
	"testing"
)

const first = 1
const second = '2'
const third = "3"
const forth = 4.00
const fifth = true

var firstList, secondList, thirdList, forthList, fifthList *LinkedList

func makeList(t *testing.T) {
	var err error
	firstList = NewLinkedList(first)

	secondList, err = firstList.Insert(second, 0)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	forthList, err = secondList.Insert(forth, 0)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	thirdList, err = firstList.Insert(third, 1)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	fifthList, err = secondList.Insert(fifth, 2)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	_, err = firstList.Insert(false, 5)
	if err == nil {
		t.Fatal("expected error")
	}
}

func tearDown() {
	firstList, secondList, thirdList, forthList, fifthList = nil, nil, nil, nil, nil
}

func TestLinkedList_InsertAndOrderElement(t *testing.T) {
	makeList(t)

	if firstList.Element != first {
		t.Error("expected element:", first, ", but get:", firstList.Element)
	}
	if secondList.Element != second {
		t.Error("expected element:", second, ", but get:", secondList.Element)
	}
	if thirdList.Element != third {
		t.Error("expected element:", third, ", but get:", thirdList.Element)
	}
	if forthList.Element != forth {
		t.Error("expected element:", forth, ", but get:", forthList.Element)
	}
	if fifthList.Element != fifth {
		t.Error("expected element:", fifth, ", but get:", fifthList.Element)
	}

	secondListCheck, err := firstList.Next()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if secondListCheck != secondList {
		t.Error("expected next:", secondList, ", but get", secondListCheck)
	}
	thirdListCheck, err := secondList.Next()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if thirdListCheck != thirdList {
		t.Error("expected next:", thirdList, ", but get", thirdListCheck)
	}
	forthListCheck, err := thirdList.Next()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if forthListCheck != forthList {
		t.Error("expected next:", forth, ", but get", forthListCheck)
	}
	fifthListCheck, err := forthList.Next()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if forthListCheck != forthList {
		t.Error("expected next:", fifthList, ", but get", fifthListCheck)
	}
	_, err = fifthList.Next()
	if err == nil {
		t.Fatal("expected error")
	}

	if firstList.Len() != 5 {
		t.Errorf("expected len: %d, but get: %d", 5, firstList.Len())
	}
	if secondList.Len() != 4 {
		t.Errorf("expected len: %d, but get: %d", 4, secondList.Len())
	}
	if thirdList.Len() != 3 {
		t.Errorf("expected len: %d, but get: %d", 3, thirdList.Len())
	}
	if forthList.Len() != 2 {
		t.Errorf("expected len: %d, but get: %d", 2, forthList.Len())
	}
	if fifthList.Len() != 1 {
		t.Errorf("expected len: %d, but get: %d", 1, fifthList.Len())
	}

	tearDown()
}

func TestLinkedList_Remove(t *testing.T) {
	makeList(t)

	_, err := firstList.Remove(5)
	if err == nil {
		t.Fatal("expected error")
	}
	_, err = firstList.Remove(0)
	if err == nil {
		t.Fatal("expected error")
	}
	if firstList.Len() != 5 {
		t.Errorf("expected len: %d, but get: %d", 5, firstList.Len())
	}
	secondListCheck, err := firstList.Remove(1)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if secondListCheck != secondList.Element {
		t.Error("expect removed element:", secondList.Element, ", but get:", secondListCheck)
	}
	if firstList.Len() != 4 {
		t.Errorf("expected len: %d, but get: %d", 5, firstList.Len())
	}
	fifthListCheck, err := firstList.Remove(3)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if fifthListCheck != fifthList.Element {
		t.Error("expect removed element:", fifthList.Element, ", but get:", fifthListCheck)
	}
	if firstList.Len() != 3 {
		t.Errorf("expected len: %d, but get: %d", 5, firstList.Len())
	}
	forthListCheck, err := thirdList.Remove(1)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if forthListCheck != forthList.Element {
		t.Error("expect removed element:", forthList.Element, ", but get:", forthListCheck)
	}
	if firstList.Len() != 2 {
		t.Errorf("expected len: %d, but get: %d", 5, firstList.Len())
	}
	thirdListCheck, err := firstList.Remove(1)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if thirdListCheck != thirdList.Element {
		t.Error("expect removed element:", thirdList.Element, ", but get:", thirdListCheck)
	}
	if firstList.Len() != 1 {
		t.Errorf("expected len: %d, but get: %d", 5, firstList.Len())
	}
	_, err = firstList.Remove(1)
	if err == nil {
		t.Fatal("expected error")
	}
	if firstList.Len() != 1 {
		t.Errorf("expected len: %d, but get: %d", 5, firstList.Len())
	}
	tearDown()
}
