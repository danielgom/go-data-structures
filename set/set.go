package main

import (
	"errors"
	"fmt"
	"log"
)

type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	return &Set{Elements: make(map[string]struct{})}
}

// Add creates a new element and places it into the set
// returns false if the element to add already existed on the set
func (s *Set) Add(elem string) bool {

	if _, exists := s.Elements[elem]; exists {
		return false
	}
	s.Elements[elem] = struct{}{}
	return true
}

// Remove removes and element from our set if it exists, otherwise it throws an error
func (s *Set) Remove(elem string) error {

	if _, exists := s.Elements[elem]; !exists {
		return errors.New("element is not present in the set")
	}
	delete(s.Elements, elem)
	return nil

}

// Contains looks for the item on the set
// returns true if the item was found, false otherwise
func (s *Set) Contains(elem string) bool {

	_, exists := s.Elements[elem]
	return exists
}

func (s *Set) List() {
	for k := range s.Elements {
		fmt.Println(k)
	}
}

// AddAll adds all elements in the slice provided
// returns true if the state of the set changes (any addition occurs)
func (s *Set) AddAll(elements []string) bool{

	modified := false
	for _, element := range elements {
		if s.Add(element) {
			modified = true
		}
	}
	return modified
}

func main() {

	s := NewSet()

	s.Add("h")
	s.Add("o")
	s.Add("l")
	s.Add("a")
	s.List()
	err := s.Remove("h")
	if err != nil {
		log.Fatalln("error thrown,", err.Error())
	}
	fmt.Println()
	s.List()
	fmt.Println(s.Contains("h"))
	s.AddAll([]string{"s", "i", "d", "n"})
	s.List()

}
