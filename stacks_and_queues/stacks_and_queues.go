package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// stack represents the queue and the stack
type stack struct {
	items []int
}

// Push will add a value at the end just like linked lists
func (s *stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove the last value added to the stack LIFO
func (s *stack) Pop() int {

	if s.items == nil {
		return 0
	}

	l := len(s.items) - 1
	rem := s.items[l]
	s.items = s.items[:l]
	return rem
}

// Dequeue will remove the first element added to the stack/queue FIFO
func (s *stack) Dequeue() int {
	if s.items == nil {
		return 0
	}

	rem := s.items[0]
	s.items = s.items[1:]
	return rem
}

func main() {

	s := &stack{}
	fmt.Println(s)
	s.Push(100)
	s.Push(200)
	s.Push(300)
	s.Push(400)
	s.Push(500)
	fmt.Println(s)

	s.Pop()
	fmt.Println(s)
	s.Dequeue()
	fmt.Println(s)

	fmt.Println(os.Hostname())
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range interfaces {
		addrs, _ := v.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Println(ip)
		}
	}
}
