package structure

import (
	"fmt"
	"time"
)

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

func (d *Deadline) Overdue() bool {
	return d != nil && (*d).Time.Before(time.Now())
}

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

func (t Task) Overdue() bool {
	return t.Deadline.Overdue()
}

func Example_taskTestAll() {
	d1 := NewDeadline(time.Now().Add(-4 * time.Hour))
	d2 := NewDeadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, d1}
	t2 := Task{"4h later", TODO, d2}
	t3 := Task{"no deadline", TODO, nil}
	fmt.Println(t1.Overdue())
	fmt.Println(t2.Overdue())
	fmt.Println(t3.Overdue())
	// Output:
	// true
	// false
	// false
}

type Address struct {
	City  string
	State string
}

type Telephone struct {
	Mobile string
	Direct string
}

type Contact struct {
	Address
	Telephone
}

func ExampleContact() {
	var c Contact
	c.Mobile = "123-456-789"
	fmt.Println(c.Telephone.Mobile)
	c.Address.City = "San Francisco"
	c.State = "CA"
	c.Direct = "N/A"
	fmt.Println(c)
	// Output:
	// 123-456-789
	// {{San Francisco CA} {123-456-789 N/A}}
}
