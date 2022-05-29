from collections import defaultdict, namedtuple
import sys
import re


Point = namedtuple("Point", "x y")


class UnionFind:
    def __init__(self):
        self.points_union = {}

    def add_point(self, point):
        if point not in self.points_union:
            self.points_union[point] = point

    def find(self, point):
        p = point
        while p != self.points_union[p]:
            p = self.points_union[p]
        return p

    def connect(self, point_a, point_b):
        parent_a = self.find(point_a)
        parent_b = self.find(point_b)
        self.points_union[parent_b] = parent_a

    def get_sets(self):
        parents_sets = defaultdict(list)
        for k in self.points_union:
            parents_sets[self.find(k)].append(k)
        return [parents_sets[k] for k in parents_sets]


def load_points():
    points_union = UnionFind()
    for line in sys.stdin:
        nums = re.findall(r"\b\d+\b", line)
        x1, y1, x2, y2 = [int(x) for x in nums]
        a, b = Point(x1, y1), Point(x2, y2)
        points_union.add_point(a)
        points_union.add_point(b)
        points_union.connect(a, b)
    return  points_union


def main():
    uf = load_points()


if __name__ == "__main__":
    main()
