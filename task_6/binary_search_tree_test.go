package task6

import "testing"

func TestBinarySearchTree(t *testing.T) {
	var bst BinarySearchTree[int, int]

	bst.Add(4, 4)
	bst.Add(2, 2)
	bst.Add(1, 1)
	bst.Add(3, 3)
	bst.Add(6, 6)
	bst.Add(5, 5)
	bst.Add(7, 7)

	if got := bst.Remove(8); got != nil {
		t.Fatalf("Remove(8) = %v, want nil", *got)
	}

	if got := bst.Get(3); got == nil || *got != 3 {
		t.Fatalf("Get(3) = %v, want 3", formatPtr(got))
	}

	if got := bst.Remove(4); got == nil || *got != 4 {
		t.Fatalf("Remove(4) = %v, want 4", formatPtr(got))
	}

	if got := bst.Get(4); got != nil {
		t.Fatalf("Get(4) = %v, want nil", *got)
	}

	if got := bst.Get(5); got == nil || *got != 5 {
		t.Fatalf("Get(5) = %v, want 5", formatPtr(got))
	}

	bst.Add(5, 10)
	if got := bst.Get(5); got == nil || *got != 10 {
		t.Fatalf("Get(5) after overwrite = %v, want 10", formatPtr(got))
	}

	if got := bst.Remove(5); got == nil || *got != 10 {
		t.Fatalf("Remove(5) = %v, want 10", formatPtr(got))
	}

	if got := bst.Get(5); got != nil {
		t.Fatalf("Get(5) after remove = %v, want nil", *got)
	}
}

func formatPtr(v *int) any {
	if v == nil {
		return nil
	}
	return *v
}
