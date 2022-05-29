package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

type UnionFind struct {
	parents map[uint64]map[uint64]*Point
}

func (uf *UnionFind) AddNewPoint(p *Point) {
	if _, okX := uf.parents[p.x]; okX {
		if _, okY := uf.parents[p.x][p.y]; okY {
			// the point already exists
			return
		}
	} else {
		uf.parents[p.x] = make(map[uint64]*Point)
	}
	uf.parents[p.x][p.y] = p
}

func (uf *UnionFind) getFromParents(p *Point) *Point {
	if _, okX := uf.parents[p.x]; okX {
		if _, okY := uf.parents[p.x][p.y]; okY {
			return uf.parents[p.x][p.y]
		}
	}
	return nil
}

func (uf *UnionFind) FindParent(p *Point) *Point {
	current := p
	parent := uf.getFromParents(current)
	if parent == nil {
		return nil
	}
	for !current.Equals(parent) {
		current = parent
	}
	return parent
}

func (uf *UnionFind) Connect(p1 *Point, p2 *Point) {
	parentP1 := uf.FindParent(p1)
	parentP2 := uf.FindParent(p2)
	if parentP1 == nil || parentP2 == nil {
		return
	}
	uf.parents[parentP2.x][parentP2.y] = parentP1
}

func load() *UnionFind {
	uf := &UnionFind{parents: make(map[uint64]map[uint64]*Point)}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		re := regexp.MustCompile(`\d+`)
		tmp := re.FindAllString(s, -1)
		if len(tmp) != 4 {
			_, _ = fmt.Fprintf(os.Stderr, "error: line %s is in incorrect format", s)
			os.Exit(-1)
		}
		var x1, y1, x2, y2 uint64
		x1, _ = strconv.ParseUint(tmp[0], 10, 64)
		y1, _ = strconv.ParseUint(tmp[1], 10, 64)
		x2, _ = strconv.ParseUint(tmp[2], 10, 64)
		y2, _ = strconv.ParseUint(tmp[3], 10, 64)
		p1 := &Point{x: x1, y: y1}
		p2 := &Point{x: x2, y: y2}
		fmt.Printf("Loaded points P1=[%d, %d], P2=[%d, %d]\n", p1.x, p2.y, p2.x, p2.y)
		uf.AddNewPoint(p1)
		uf.AddNewPoint(p2)
		uf.Connect(p1, p2)
	}
	return uf
}

func main() {
	_ = load()
}
