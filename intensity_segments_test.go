package main

import (
	"testing"
)

func TestIntensitySegments(t *testing.T) {
	//segments := NewIntensitySegments()
	//
	//// Test initial state
	//if got := segments.ToString(); got != "[]" {
	//	t.Errorf("expected '[]', got %s", got)
	//}
	//
	//// Test Add method
	//segments.Add(10, 30, 1)
	//if got := segments.ToString(); got != "[[10,1],[30,0]]" {
	//	t.Errorf("expected '[[10,1],[30,0]]', got %s", got)
	//}
	//
	//segments.Add(20, 40, 1)
	//if got := segments.ToString(); got != "[[10,1],[20,2],[30,1],[40,0]]" {
	//	t.Errorf("expected '[[10,1],[20,2],[30,1],[40,0]]', got %s", got)
	//}
	//
	//segments.Add(10, 40, -2)
	//if got := segments.ToString(); got != "[[10,-1],[20,0],[30,-1],[40,0]]" {
	//	t.Errorf("expected '[[10,-1],[20,0],[30,-1],[40,0]]', got %s", got)
	//}
	//
	//segments.Add(10, 40, -1)
	//if got := segments.ToString(); got != "[[10,-2],[20,-1],[30,-2],[40,0]]" {
	//	t.Errorf("expected '[[10,-2],[20,-1],[30,-2],[40,0]]', got %s", got)
	//}
	//
	//segments.Add(10, 40, -1)
	//if got := segments.ToString(); got != "[[10,-3],[20,-2],[30,-3],[40,0]]" {
	//	t.Errorf("expected '[[10,-3],[20,-2],[30,-3],[40,0]]', got %s", got)
	//}
	//
	//// Reset and test another scenario
	//segments = NewIntensitySegments()
	//if got := segments.ToString(); got != "[]" {
	//	t.Errorf("expected '[]', got %s", got)
	//}
	//
	//segments.Add(10, 30, 1)
	//if got := segments.ToString(); got != "[[10,1],[30,0]]" {
	//	t.Errorf("expected '[[10,1],[30,0]]', got %s", got)
	//}
	//
	//segments.Add(20, 40, 1)
	//if got := segments.ToString(); got != "[[10,1],[20,2],[30,1],[40,0]]" {
	//	t.Errorf("expected '[[10,1],[20,2],[30,1],[40,0]]', got %s", got)
	//}
	//
	//segments.Add(10, 40, -1)
	//if got := segments.ToString(); got != "[[20,1],[30,0]]" {
	//	t.Errorf("expected '[[20,1],[30,0]]', got %s", got)
	//}
	//
	//segments.Add(10, 40, -1)
	//if got := segments.ToString(); got != "[[10,-1],[20,0],[30,-1],[40,0]]" {
	//	t.Errorf("expected '[[10,-1],[20,0],[30,-1],[40,0]]', got %s", got)
	//}

	// Reset and test set scenario
	segments := NewIntensitySegments()
	if got := segments.ToString(); got != "[]" {
		t.Errorf("expected '[]', got %s", got)
	}

	segments.Set(10, 30, 1)
	if got := segments.ToString(); got != "[[10,1],[30,0]]" {
		t.Errorf("expected '[[10,1],[30,0]]', got %s", got)
	}

	segments.Set(20, 40, 1)
	if got := segments.ToString(); got != "[[10,1],[40,0]]" {
		t.Errorf("expected '[[10,1],[40,0]]', got %s", got)
	}

	segments.Set(10, 30, -1)
	if got := segments.ToString(); got != "[[10,-1],[30,1],[40,0]]" {
		t.Errorf("expected '[[10,-1],[30,1],[40,0]]', got %s", got)
	}

	segments.Set(10, 15, -2)
	if got := segments.ToString(); got != "[[10,-2],[15,-1],[30,1],[40,0]]" {
		t.Errorf("expected '[[10,-2],[[15,-1],[30,1],[40,0]]', got %s", got)
	}

}
