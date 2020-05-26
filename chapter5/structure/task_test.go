package structure

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
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
	Title    string    `json:"title"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

func (t Task) Overdue() bool {
	return t.Deadline.Overdue()
}

func Example_taskTestAll() {
	d1 := NewDeadline(time.Now().Add(-4 * time.Hour))
	d2 := NewDeadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, d1, 0, nil}
	t2 := Task{"4h later", TODO, d2, 0, nil}
	t3 := Task{"no deadline", TODO, nil, 0, nil}
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

func (s status) MarshalJSON() ([]byte, error) {
	switch s {
	case UNKNOWN:
		return []byte(`"UNKNOWN"`), nil
	case TODO:
		return []byte(`"TODO"`), nil
	case DONE:
		return []byte(`"DONE"`), nil
	default:
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
}

func (s *status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"UNKNOWN"`:
		*s = UNKNOWN
	case `"TODO"`:
		*s = TODO
	case `"DONE"`:
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}

func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(unix, 0)
	return nil
}

func Example_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
		0,
		nil,
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// {"title":"Laundry","status":"DONE","deadline":1439739780}
}

func Example_unmarshalJSON() {
	b := []byte(`{"title":"Buy Milk","status":"DONE","deadline":1439739780}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
	// Output:
	// Buy Milk
	// 2
	// 2015-08-16 15:43:00 +0000 UTC
}

func Example_mapMarshalJSON() {
	b, _ := json.Marshal(map[string]interface{}{
		"Name": "John",
		"Age":  16,
	})
	fmt.Println(string(b))
	// Output:
	// {"Age":16,"Name":"John"}
}

type Fields struct {
	VisibleField   string `json:"visibleField"`
	InvisibleField string `json:"invisibleField"`
}

func ExampleOmitFields() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(struct {
		*Fields
		InvisibleField string `json:"invisibleField,omitempty"`
		Additional     string `json:"additional,omitempty"`
	}{Fields: f, Additional: "c"})
	fmt.Println(string(b))
	// Output:
	// {"visibleField":"a","additional":"c"}
}

func Example_gob() {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	data := map[string]string{"N": "J"}
	if err := enc.Encode(data); err != nil {
		fmt.Println(err)
	}
	const width = 16
	for start := 0; start < len(b.Bytes()); start += width {
		end := start + width
		if end > len(b.Bytes()) {
			end = len(b.Bytes())
		}
		fmt.Printf("% x\n", b.Bytes()[start:end])
	}
	dec := gob.NewDecoder(&b)
	var restored map[string]string
	if err := dec.Decode(&restored); err != nil {
		fmt.Println(err)
	}
	fmt.Println(restored)
	// Output:
	// 0e ff 81 04 01 02 ff 82 00 01 0c 01 0c 00 00 08
	// ff 82 00 01 01 4e 01 4a
	// map[N:J]
}

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

func Example_PrintStringer() {
	t := Task{
		"Laundry",
		DONE,
		nil,
		0,
		nil,
	}
	fmt.Print(t)
	// Output:
	// [v] Laundry <nil>
}

type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix+"\t")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}

func ExampleIncludeSubTasks_String() {
	fmt.Println(IncludeSubTasks(Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{
			{
				Title:    "Wash",
				Status:   TODO,
				Deadline: nil,
				Priority: 2,
				SubTasks: []Task{
					{"Put", DONE, nil, 2, nil},
					{"Detergent", TODO, nil, 2, nil},
				},
			}, {
				Title:    "Dry",
				Status:   TODO,
				Deadline: nil,
				Priority: 2,
				SubTasks: nil,
			}, {
				Title:    "Fold",
				Status:   TODO,
				Deadline: nil,
				Priority: 2,
				SubTasks: nil,
			},
		},
	}))
	// Output:
	// [ ] Laundry <nil>
	//	[ ] Wash <nil>
	//		[v] Put <nil>
	//		[ ] Detergent <nil>
	//	[ ] Dry <nil>
	//	[ ] Fold <nil>
}
