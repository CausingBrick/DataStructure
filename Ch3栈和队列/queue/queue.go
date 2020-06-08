package main

import "fmt"

// QueueArr implements queue by slice.
type QueueArr []interface{}

// Enqueue push element to the tail of queue.
func (q *QueueArr) Enqueue(x ...interface{}) {
	for _, val := range x {
		*q = append(*q, val)
	}
}

// Init a queue of length 0 and capacity 10.
func (q *QueueArr) Init() QueueArr {
	*q = make([]interface{}, 0, 10)
	return *q
}

// Dequeue remove the element from the head.
func (q *QueueArr) Dequeue() interface{} {
	queue := *q
	if len(queue) == 0 {
		return nil
	}
	*q = queue[1:]
	return q
}

// Search return the index of val in queue or -1 if not found.
func (q *QueueArr) Search(val interface{}) int {
	if len(*q) == 0 {
		return -1
	}
	for k, v := range *q {
		if v == val {
			return k
		}
	}
	return -1
}

// IsEmpty return if the queue is empty.
func (q QueueArr) IsEmpty() bool {
	return len(q) == 0
}

// Show print the queue.
func (q QueueArr) Show() {
	for _, val := range q {
		fmt.Printf("%d <-", val)
	}
	fmt.Println()
}

// Element is the element of queue.
type Element struct {
	next, prev *Element
	Value      interface{}
}

// Queue represents a queue implement by double linked list.
type Queue struct {
	// The head stores the tail with head.next and top with head.prev.
	head Element
	// Current queue lenght exclude (this) sintinel element.
	len int
}

// Init initializes queue or clears s.
func (s *Queue) Init() *Queue {
	s.head.next = &s.head
	s.head.prev = &s.head
	s.len = 0
	return s
}

// Enqueue adds a element to the tail.
func (s *Queue) Enqueue(v ...interface{}) {
	for i := 0; i < len(v); i++ {
		e := &Element{s.head.next, &s.head, v[i]}
		s.head.next.prev = e
		s.head.next = e
		s.len++
	}
}

// Dequeue removes the element at head.
func (s *Queue) Dequeue() *Element {
	if s.len == 0 {
		return nil
	}
	current := s.head.prev
	s.head.prev = current.prev
	s.head.prev.next = &s.head
	current.prev, current.next = nil, nil
	s.len--
	return current
}

// Show prints all the element's value.
func (s *Queue) Show() {
	if s.len == 0 {
		return
	}
	for e := s.head.next; e != &s.head; e = e.next {
		fmt.Printf("-> %v", e.Value)
	}
	fmt.Println()
}

// IsEmpty return if the queue is empty.
func (s *Queue) IsEmpty() bool {
	return s.len == 0
}

// Search returns element stroed val or nil if not exist.
func (s *Queue) Search(val interface{}) *Element {
	if s.len == 0 {
		return nil
	}
	for e := s.head.next; e != &s.head; e = e.next {
		if e.Value == val {
			return e
		}
	}
	return nil
}
func main() {
	fmt.Println("Test queue impelement by array")
	q1 := new(QueueArr).Init()
	q1.Enqueue(1, 2, 3, 4, 5, 6, 7, 8)
	q1.Show() //1 <-2 <-3 <-4 <-5 <-6 <-7 <-8 <-
	q1.Dequeue()
	q1.Show()                                        //1 <-2 <-3 <-4 <-5 <-6 <-7 <-8 <-
	fmt.Println("The index of 3 is:", q1.Search(3))  //The index of 3 is: 1
	fmt.Println("Is the queue empty?", q1.IsEmpty()) //Is the queue empty? false
	q1.Init()
	fmt.Println("Is the queue empty?", q1.IsEmpty()) //Is the queue empty? true
	fmt.Println("Test queue impelement by list")
	q2 := new(Queue).Init()
	q2.Enqueue(1, 2, 3, 4, 5, 6, 7, 8)
	q2.Show() //-> 8-> 7-> 6-> 5-> 4-> 3-> 2-> 1
	q2.Dequeue()
	q2.Show()                                         //-> 8-> 7-> 6-> 5-> 4-> 3-> 2
	fmt.Println("The element of 3 is:", q2.Search(3)) //The element of 3 is: &{0xc0000ae040 0xc0000ae080 3}
	fmt.Println("Is the queue empty?", q2.IsEmpty())  //Is the queue empty? false
	q2.Init()
	fmt.Println("Is the queue empty?", q2.IsEmpty()) //Is the queue empty? true
}
