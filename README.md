# Union find algorithm
In this repository, you will find multiple implementations of the union find algorithm for points in a 2D plane.

All implementations in this repository do the following:
- read input line by line from `stdin` until `EOF`, where each line contains two points which are connected. Each line is in format `[x1, y1] [x2, y2]`, e. g. `[2, 3] [5, 4]`. This means that points `P1 = [2, 3]` and `P2 = [5, 4]` are connected with an edge in an undirected graph.
- construct the union find data structure. This data structure represents connected components in the loaded graph.

For instance, consider the following input:
```
[1, 3] [2, 5]
[3, 4] [1, 3]
[8, 6] [6, 8]
```
This input represents an undirected graph with:
- 5 vertices ([1, 3], [2, 5], [3, 4], [6, 8], [8, 6])
- 2 connected components:
	- [1, 3], [2, 5], [3, 4]
	- [6, 8], [8, 6]

Note, that lines `[1, 3] [2, 5]` and `[2, 5] [1, 3]` are equivalent. Also note, that spacing of the lines does not play any role. Implementations must correctly load all of the following lines:
- `[1, 3] [2, 5]`
- `[1,3][2,5]`
- `[1,3] [2,5]`
- `[1, 3] [2,5]`
- etc.

The input is guaranteed to be in the corre
## Why
Now would be a good time to talk about the why of this project. The **sole** purpose of this project is for me to learn basics of several interesting languages and create a sort of Rosetta stone that I can refer back to in the future.

Why select union find for this reason though? Well, my reasoning is that in order to implement the union find algorithm for the problem described above, one must know the following:
- how to use standard library in the given language, notably HashMap (in case the given language has such feature)
- how to construct abstract data types. This also implies learning about memory management of the given language.
- how to handle IO and parsing. This also implies learning about error handling mechanisms of the given language.

If you want a "base" implementation, refer to the python version. I wrote it as a part of a school assignment. The other versions more or less emulate the python one.
## Disclaimer
The implementations you will find here are not very efficient. Furthermore, I am not an expert on most of the used languages. Therefore, I cannot guarantee that the implementations are idiomatic.