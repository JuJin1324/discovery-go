package graph

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

type Graph [][]int

func (adjList Graph) WriteTo(w io.Writer) error {
	size := len(adjList)
	if _, err := fmt.Fprintf(w, "%d", size); err != nil {
		return err
	}
	for i := 0; i < size; i++ {
		lsize := len(adjList[i])
		if _, err := fmt.Fprintf(w, "\n%d", lsize); err != nil {
			return err
		}
		for j := 0; j < lsize; j++ {
			if _, err := fmt.Fprintf(w, " %d", adjList[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
	return nil
}

func (adjList *Graph) ReadFrom(r io.Reader) error {
	var size int
	if _, err := fmt.Fscanf(r, "%d", &size); err != nil {
		return err
	}
	*adjList = make([][]int, size)
	for i := 0; i < size; i++ {
		var lsize int
		if _, err := fmt.Fscanf(r, "\n%d", &lsize); err != nil {
			return err
		}
		(*adjList)[i] = make([]int, lsize)
		for j := 0; j < lsize; j++ {
			if _, err := fmt.Fscanf(r, " %d", &(*adjList)[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fscanf(r, "\n"); err != nil {
		return err
	}
	return nil
}

func TestWriteTo(t *testing.T) {
	adjList := Graph{
		{3, 4},
		{0, 2},
		{3},
		{2, 4},
		{0},
	}
	w := bytes.NewBuffer(nil)
	if err := adjList.WriteTo(w); err != nil {
		t.Error(err)
	}
	expected := "5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n"
	if expected != w.String() {
		t.Logf("expected: %s\n", expected)
		t.Errorf("found: %s\n", w.String())
	}
}

func ExampleReadFrom() {
	r := strings.NewReader("5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n")
	var adjList Graph
	if err := (&adjList).ReadFrom(r); err != nil {
		fmt.Println(err)
	}
	fmt.Println(adjList)
	// Output:
	// [[3 4] [0 2] [3] [2 4] [0]]
}
