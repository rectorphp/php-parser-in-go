package ast

import "testing"

func TestAttributeBagSetGet(test *testing.T) {
	var bag AttributeBag

	if _, ok := bag.GetAttribute("type"); ok {
		test.Fatal("empty bag reported a value")
	}

	bag.SetAttribute("type", "User")
	value, ok := bag.GetAttribute("type")
	if !ok || value != "User" {
		test.Fatalf("got %v, %v; want User, true", value, ok)
	}

	bag.SetAttribute("type", "Admin")
	if value, _ := bag.GetAttribute("type"); value != "Admin" {
		test.Fatalf("overwrite failed: got %v", value)
	}
}

func TestAttributeBagEqual(test *testing.T) {
	var empty, other AttributeBag
	if !empty.Equal(other) {
		test.Fatal("two empty bags should be equal")
	}

	set := AttributeBag{}
	set.SetAttribute("type", "User")
	if empty.Equal(set) {
		test.Fatal("empty bag should differ from a populated one")
	}

	same := AttributeBag{}
	same.SetAttribute("type", "User")
	if !set.Equal(same) {
		test.Fatal("bags with the same attributes should be equal")
	}
}

// node embeds AttributeBag, so every node satisfies the Vertex attribute methods.
func TestNodeCarriesAttributes(test *testing.T) {
	var node Vertex = &ExprVariable{}
	node.SetAttribute("type", "string")
	if value, ok := node.GetAttribute("type"); !ok || value != "string" {
		test.Fatalf("node attribute not stored: got %v, %v", value, ok)
	}
}
