import gleam/int
import gleam/list

pub fn convert(number: Int) -> String {
  let c = [#(3, "Pling"), #(5, "Plang"), #(7, "Plong")]
  c
  |> list.fold("", fn(acc, tup) {
    let #(n, s) = tup
    case number % n == 0 {
      True -> acc <> s
      _ -> acc
    }
  })
  |> fn(s) {
    case s {
      "" -> int.to_string(number)
      _ -> s
    }
  }
}
