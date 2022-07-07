# Notes on the Go implementation
First, I must conclude, that working with go was a pleasant experience. The tooling is amazing, the standard library is very solid, docs are informative, etc.

Nonetheless, there were some hurdles:
- I do not like, that `go mod` command de facto assumes, that you have a github/gitlab account. Sure, I know, that I could simply use a dummy address, e.g. `go mod init example.com/project` or simply type `go mod init project`, but I wouldn't know that if I relied solely on the docs. Some people also complain about the package management compared to the likes of Rust's `cargo`, though I cannot testify to this end.
- There is no way to use a custom key in a `map`. Originally, I wanted to emulate the python version, which relies on a dictionary with tuples as both keys and values. This is not possible in Go. I do understand that there are workarounds. I also do understand the motivation for this design choice (and dare I say, I kind of approve of it), but this is still a limitation, that I did not expect to run into. In the nest section, I will explain, how I dealt with it.
- I would appreciate a function to test for existence of a key in map. Sometimes, all I need is to test whether a key exists and `if _, ok := m[key]; ok { ... }` feels needlessly verbose. That said, this is such a minor complaint, I did consider not to include it in this writeup.
- This thing:
	- `val, ok := m[key]`, makes sense
	- `val, _ := m[key]`, fine, very pythonic
	- `_, ok := m[key]`, fine, very pythonic
	- `_, _ = m[key]`, why the missing colon? Ok, I get why. No new variable is created. That's what the error says. But this still feels very inconsistent. The thing is, a variable *is* created, two of those, in fact. They are created by the callee, but they are *ignored* by the caller. Instead, Go pretends there is nothing to see and no variables created. Once again, this is a minor thing. I just think that semantically, it's a strange way to look at the ignored values.

## Map workaround
I had several ideas to work around the map keys issue:
1) write a hash function that takes in a point and returns a `uint64` hash of that point. Use this value as the key. This solution would be perfect for 99% of the use cases. However...
2) ... I had the "brilliant" idea to instead represent the points as a map of maps of point. Each point has `x` and `y` coordinates. The `x` coordinate would be the key in the parents map. The value would be a map of `y` coordinates and pointers to points.
3) Since I already needed to print points a readable way and implement a `ToString()` method on that struct, I decided, I might as well use that string as the key to the parents map. This is a better solution than 1. because it saves me from writing a custom hash function (though that wouldn't be too difficult). This is the solution I used in the final version. It suffers from the same problems as 1., though.

The main problem of 1. and 3. is that a part of the information about the keys is lost. In this particular use case, it is a problem.

The parents map should contain key-value pairs with the semantics of point-parent. After the union find data structure is constructed, chances are one would like to use it. But for that you need to know the points. While one could reconstruct the points from string representation used as keys (in the case of 3.), this wouldn't be efficient and also wouldn't work for a more complex ADT. I solved this by adding another map, which stores the points. This is fast, but obviously requires more memory.

The strategy 2. also struggles with this aspect, though if one only needed to know the coordinates, this would be a good enough solution. I initially implemented this strategy, but the resulting code, while functional, was needlessly complicated.

Therefore, I settled for 3.