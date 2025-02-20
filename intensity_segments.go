package main

import (
	"encoding/json"
	"fmt"
)

// Node represents a segment in the linked list.
type Node struct {
	point     int
	intensity int
	next      *Node
}

// IntensitySegments manages intensity segments using a linked list.
type IntensitySegments struct {
	head *Node
}

// Add adds intensity to the segment defined by the range [from, to).
func (is *IntensitySegments) Add(from, to, amount int) {
	is.updateSegments(from, to, amount, true)
}

// Set sets the intensity for the segment defined by the range [from, to).
func (is *IntensitySegments) Set(from, to, amount int) {
	is.updateSegments(from, to, amount, false)
}

// updateSegments updates the segments based on the provided range and intensity.
func (is *IntensitySegments) updateSegments(from, to, amount int, isAdd bool) {
	// Helper to find or insert a node in the linked list
	var prev *Node
	current := is.head

	for current != nil && current.point < from {
		prev = current
		current = current.next
	}

	// Insert or update the 'from' point
	if current == nil || current.point != from {
		newNode := &Node{point: from, intensity: 0, next: current}
		if prev != nil {
			newNode.intensity = prev.intensity
			prev.next = newNode
		} else {
			is.head = newNode
		}
		current = newNode
		//current = newNode.next
	}

	// Update the intensity from 'from' to 'to'
	for current != nil && current.point < to {
		if isAdd {
			current.intensity += amount
		} else {
			current.intensity = amount
		}

		// Move to the next segment
		prev = current
		current = current.next
	}

	if current == nil || current.point != to {
		// Insert or update the 'to' point
		newNode := &Node{point: to, intensity: 0, next: current}
		if prev != nil {
			prev.next = newNode
		} else {
			is.head = newNode
		}
	}

	// Clean up: remove segments with same intensity
	is.cleanUp()
}

// cleanUp removes segments with the same intensity as the previous segment.
func (is *IntensitySegments) cleanUp() {
	current := is.head
	for current == is.head && current != nil && current.intensity == 0 {
		is.head = current.next
		current = current.next
	}
	for current != nil && current.next != nil {
		if current.intensity == current.next.intensity {
			// Skip the next node
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
}

// ToString returns a string representation of the current segments.
func (is *IntensitySegments) ToString() string {
	var segments [][]int
	current := is.head
	for current != nil {
		segments = append(segments, []int{current.point, current.intensity})
		current = current.next
	}
	if len(segments) == 0 {
		return "[]"
	}
	jsonData, _ := json.Marshal(segments)
	return string(jsonData)
}

// Example usage
func main() {
	segments := &IntensitySegments{}
	fmt.Println(segments.ToString()) // Should be "[]"
	segments.Add(10, 30, 1)
	fmt.Println(segments.ToString()) // Should be: "[[10,1],[30,0]]"
	segments.Add(20, 40, 1)
	fmt.Println(segments.ToString()) // Should be: "[[10,1],[20,2],[30,1],[40,0]]"
	segments.Add(10, 40, -2)
	fmt.Println(segments.ToString()) // Should be: "[[10,-1],[20,0],[30,-1],[40,0]]"
	segments.Add(10, 40, -1)
	fmt.Println(segments.ToString()) // Should be "[[10,-2],[20,0],[30,-2], [40,0]]"
	segments.Add(10, 40, -1)
	segments = &IntensitySegments{}
	fmt.Println(segments.ToString()) // Should be "[]"
	segments.Add(10, 30, 1)
	fmt.Println(segments.ToString()) // Should be "[[10,1],[30,0]]"
	segments.Add(20, 40, 1)
	fmt.Println(segments.ToString()) // Should be "[[10,1],[20,2],[30,1],[40,0]]"
	segments.Add(10, 40, -1)
	fmt.Println(segments.ToString()) // Should be "[[20,1],[30,0]]"
	segments.Add(10, 40, -1)
	fmt.Println(segments.ToString()) // Should be "[[10,-1],[20,0],[30,-1],[40,0]]"
}
