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
	pre       *Node
}

// IntensitySegments manages intensity segments using a linked list.
type IntensitySegments struct {
	head *Node
	tail *Node
}

func NewIntensitySegments() *IntensitySegments {
	segments := &IntensitySegments{
		head: &Node{},
		tail: &Node{},
	}
	segments.head.next = segments.tail
	segments.tail.pre = segments.head
	return segments
}

// Add adds intensity to the segment defined by the range [from, to).
func (is *IntensitySegments) Add(from, to, amount int) {
	is.addSegments(from, to, amount)
}

// Set sets the intensity for the segment defined by the range [from, to).
func (is *IntensitySegments) Set(from, to, amount int) {
	is.setSegments(from, to, amount)
}

// setSegments sets the segments based on the provided range and intensity.
func (is *IntensitySegments) setSegments(from, to, amount int) {
	// Helper to find or insert a node in the linked list
	current := is.tail.pre

	for current != is.head && current.point > to {
		current = current.pre
	}

	// Insert the 'to' point
	if current == is.head || current.point != to {
		newNode := &Node{point: from, intensity: 0, pre: current, next: current.next}
		if current.pre != is.head {
			newNode.intensity = current.pre.intensity
		}
		newNode.pre.next = newNode
		newNode.next.pre = newNode
		current = current.pre
		//current = newNode.next
	}

	// Update the intensity from 'from' to 'to'
	for current != is.head && current.point >= from {
		current.intensity = amount
		// Move to the next segment
		current = current.next
	}

	if current == is.head || current.point != from {
		// Insert the 'to' point
		newNode := &Node{point: to, intensity: amount, pre: current, next: current.next}
		newNode.pre.next = newNode
		newNode.next.pre = newNode
	}

	// Clean up: remove segments with same intensity
	is.cleanUp()
}

// addSegments adds the segments based on the provided range and intensity.
func (is *IntensitySegments) addSegments(from, to, amount int) {
	// Helper to find or insert a node in the linked list
	current := is.head.next

	for current != is.tail && current.point < from {
		current = current.next
	}

	// Insert the 'from' point
	if current == is.tail || current.point != from {
		newNode := &Node{point: from, intensity: 0, pre: current.pre, next: current}
		if current.pre != is.head {
			newNode.intensity = current.pre.intensity
		}
		newNode.pre.next = newNode
		newNode.next.pre = newNode
		current = newNode
	}

	// Update the intensity from 'from' to 'to'
	for current != is.tail && current.point < to {
		current.intensity += amount
		// Move to the next segment
		current = current.next
	}

	if current == is.tail || current.point != to {
		// Insert the 'to' point
		newNode := &Node{point: to, intensity: 0, pre: current.pre, next: current}
		if current.pre != is.head {
			newNode.intensity = current.pre.intensity
			if current.pre.point >= from && current.pre.point < to {
				newNode.intensity -= amount
			}
		}
		newNode.pre.next = newNode
		newNode.next.pre = newNode
	}

	// Clean up: remove segments with same intensity
	is.cleanUp()
}

// cleanUp removes segments with the same intensity as the previous segment.
func (is *IntensitySegments) cleanUp() {
	current := is.head.next
	for current.pre == is.head && current != is.tail && current.intensity == 0 {
		current.next.pre = current.pre
		current.pre.next = current.next
		current = current.next
	}
	for current != is.tail && current.next != is.tail {
		if current.intensity == current.next.intensity {
			// Skip the next node
			current.next = current.next.next
			current.next.pre = current
		} else {
			current = current.next
		}
	}
}

// ToString returns a string representation of the current segments.
func (is *IntensitySegments) ToString() string {
	var segments [][]int
	current := is.head.next
	for current != is.tail {
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
