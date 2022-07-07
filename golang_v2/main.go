package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y uint64
}

func (p *Point) Equals(other *Point) bool {
	if other == nil {
		return false
	} else {
		return p.x == other.x && p.y == other.y
	}
}

func (p *Point) ToString() string {
	return fmt.Sprintf("[%d, %d]", p.x, p.y)
}

type UnionFind struct {
	points  []*Point
	parents []uint64
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		points:  make([]*Point, 0, 10),
		parents: make([]uint64, 0, 10),
	}
}

func (uf *UnionFind) FindIndex(p *Point) (uint64, error) {
	for idx, point := range uf.points {
		if point.Equals(p) {
			return uint64(idx), nil
		}
	}
	return uint64(len(uf.points)), errors.New("there is no such point in the union find")
}

func (uf *UnionFind) AddNewPoint(p *Point) uint64 {
	if pos, err := uf.FindIndex(p); err == nil {
		return pos
	} else {
		pos := uint64(len(uf.points))
		uf.points = append(uf.points, p)
		uf.parents = append(uf.parents, pos)
		return pos
	}
}

func (uf *UnionFind) FindParent(p *Point) *Point {
	if current, err := uf.FindIndex(p); err != nil {
		return nil
	} else {
		parent := uf.parents[current]
		for current != parent {
			current = parent
		}
		return uf.points[parent]
	}
}

func (uf *UnionFind) Connect(p1 *Point, p2 *Point) {
	parentP1 := uf.FindParent(p1)
	parentP2 := uf.FindParent(p2)
	if parentP1 == nil || parentP2 == nil || parentP1.Equals(parentP2) {
		return
	}
	pos1, err1 := uf.FindIndex(parentP1)
	pos2, err2 := uf.FindIndex(parentP2)
	if err1 != nil || err2 != nil {
		return
	}
	uf.parents[pos2] = pos1
}

func (uf *UnionFind) GetConnectedComponents() map[string][]*Point {
	result := make(map[string][]*Point)
	for key, value := range uf.points {
		parent := uf.FindParent(value)
		parentStr := parent.ToString()
		if v, ok := result[parentStr]; ok {
			result[parentStr] = append(v, uf.points[key])
		} else {
			result[parentStr] = []*Point{uf.points[key]}
		}
	}
	return result
}

func load() *UnionFind {
	uf := NewUnionFind()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		// remove newlines and spaces
		s = strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), " ", "")
		if len(s) == 0 {
			// skip empty lines
			continue
		}

		// load numbers
		var x1, y1, x2, y2 uint64
		if n, err := fmt.Sscanf(s, "[%d,%d][%d,%d]", &x1, &y1, &x2, &y2); n != 4 || err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(os.Stderr, "error: line %s is in incorrect format", s)
			os.Exit(-1)
		}

		// construct points and register them in the unionfind
		p1 := &Point{x: x1, y: y1}
		p2 := &Point{x: x2, y: y2}
		uf.AddNewPoint(p1)
		uf.AddNewPoint(p2)
		uf.Connect(p1, p2)
	}
	return uf
}

func main() {
	uf := load()
	components := uf.GetConnectedComponents()
	for _, component := range components {
		for _, point := range component {
			fmt.Printf("%s ", point.ToString())
		}
		fmt.Println()
	}
}
