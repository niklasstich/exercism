pub fn equilateral(a: Float, b: Float, c: Float) -> Bool {
  a == b && b == c && valid_inequality(a, b, c)
}

pub fn isosceles(a: Float, b: Float, c: Float) -> Bool {
  { a == b || b == c || c == a } && valid_inequality(a, b, c)
}

pub fn scalene(a: Float, b: Float, c: Float) -> Bool {
  a != b && b != c && c != a && valid_inequality(a, b, c)
}

fn valid_inequality(a: Float, b: Float, c: Float) -> Bool {
  a >. 0.0
  && b >. 0.0
  && c >. 0.0
  && { a +. b >=. c && b +. c >=. a && a +. c >=. b }
}
