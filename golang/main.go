package main

import (
	"bufio"
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
	parents map[string]*Point
	points  map[string]*Point
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parents: make(map[string]*Point),
		points:  make(map[string]*Point),
	}
}

func (uf *UnionFind) AddNewPoint(p *Point) {
	if _, ok := uf.parents[p.ToString()]; ok {
		// point already in union find
		return
	} else {
		uf.parents[p.ToString()] = p
		uf.points[p.ToString()] = p
	}
}

func (uf *UnionFind) FindParent(p *Point) *Point {
	current := p
	if parent, ok := uf.parents[current.ToString()]; !ok {
		return nil
	} else {
		for !current.Equals(parent) {
			current = parent
			parent = uf.parents[parent.ToString()]
		}
		return parent
	}
}

func (uf *UnionFind) Connect(p1 *Point, p2 *Point) {
	parentP1 := uf.FindParent(p1)
	parentP2 := uf.FindParent(p2)
	if parentP1 == nil || parentP2 == nil || parentP1.Equals(parentP2) {
		return
	}
	uf.parents[parentP2.ToString()] = parentP1
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
