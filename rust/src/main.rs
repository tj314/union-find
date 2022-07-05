mod union_find {

    // this macro was shamelessly stolen from here
    // https://stackoverflow.com/a/31048103
    macro_rules! scan {
        ( $string:expr, $sep:expr, $( $x:ty ),+ ) => {{
            let mut iter = $string.split($sep);
            ($(iter.next().and_then(|word| word.parse::<$x>().ok()),)*)
        }}
    }

    use std::fmt;
    use std::io;
    use std::collections::HashMap;

    #[derive(Eq, PartialEq)]
    pub struct Point {
        x: usize,
        y: usize,
    }

    impl fmt::Display for Point {
        fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
            write!(f, "[{}, {}]", self.x, self.y)
        }
    }

    pub struct UnionFind {
        points: Vec<Point>,
        parents: Vec<usize>,
    }

    impl UnionFind {
        pub fn load() -> UnionFind {
            let mut uf = UnionFind {
                points: Vec::new(),
                parents: Vec::new(),
            };

            loop {
                let mut buff = String::new();
                match io::stdin().read_line(&mut buff) {
                    Ok(0) => break, // EOF
                    Ok(_) => {},
                    Err(_) => continue,
                }
                buff = buff.chars().map(move |c| match c {
                    c if c.is_whitespace() => String::new(),
                    '[' => String::from(","),
                    ']' => String::new(),
                    _ => c.to_string(),
                }).skip(1).collect();
                let (x1, y1, x2, y2) = scan!(buff, |c| c == ',', usize, usize, usize, usize);
                let p1 = Point{
                    x: x1.unwrap(),
                    y: y1.unwrap(),
                };
                let p2 = Point{
                    x: x2.unwrap(),
                    y: y2.unwrap(),
                };
                let pos1 = uf.add_point(p1);
                let pos2 = uf.add_point(p2);
                uf.connect(pos1, pos2);
            }
            uf
        }

        pub fn index(&self, p: &Point) -> Option<usize> {
            for (i, point) in self.points.iter().enumerate() {
                if *p == *point {
                    return Some(i);
                }
            }
            None
        }

        pub fn add_point(&mut self, p: Point) -> usize {
            if let Some(i) = self.index(&p) {
                i
            } else {
                self.points.push(p);
                let pos = self.parents.len();
                self.parents.push(self.parents.len());
                pos
            }
        }

        pub fn find_parent(&self, point: usize) -> Option<usize> {
            let mut p = point;
            loop {
                if self.parents.len() <= p {
                    return None;
                } else {
                    if self.parents[p] == p {
                        return Some(p);
                    } else {
                        p = self.parents[p];
                    }
                }
            }
        }

        pub fn connect(&mut self, p1: usize, p2: usize) {
            if let Some(parent_p1) = self.find_parent(p1) {
                if let Some(parent_p2) = self.find_parent(p2) {
                    self.parents[parent_p2] = parent_p1;
                }
            }
        }

        pub fn get_connected_components(&self) -> HashMap<usize, Vec<&Point>> {
            let mut out = HashMap::new();
            for i in 0..self.points.len() {
                let parent = self.find_parent(i).unwrap();
                let e = out.entry(parent).or_insert(Vec::new());
                e.push(&self.points[i]);
            }
            out
        }
    }
}
fn main() {
    use union_find as uf;
    let union = uf::UnionFind::load();
    let components = union.get_connected_components();
    for (_, component) in components {
        for p in component {
            print!("{} ", p);
        }
        println!();
    }
}
