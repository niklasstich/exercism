import gleam/list

pub fn sum(factors factors: List(Int), limit limit: Int) -> Int {
  factors
  |> list.map(fn(factor) {
    let rest = { limit - 1 } % factor
    let factor2 = { limit - rest } / factor
    case factor2 {
      _ if factor == 1 -> list.range(0, limit - 1)
      _ if factor2 <= 0 -> []
      _ -> list.range(0, factor2) |> list.map(fn(i) { i * factor })
    }
  })
  |> list.flatten
  |> list.unique
  |> list.fold(0, fn(acc, n) { acc + n })
}
