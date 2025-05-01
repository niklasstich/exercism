import gleam/list

const letters = [
  #(1000, "M"),
  #(900, "CM"),
  #(500, "D"),
  #(400, "CD"),
  #(100, "C"),
  #(90, "XC"),
  #(50, "L"),
  #(40, "XL"),
  #(10, "X"),
  #(9, "IX"),
  #(5, "V"),
  #(4, "IV"),
  #(1, "I"),
]

pub fn convert(number: Int) -> String {
  case number {
    _ if number > 0 ->
      letters
      |> list.fold_until(#("", number), fn(acc, tup) {
        let #(n, s) = tup
        let #(_, i) = acc
        case i / n > 0 {
          True -> list.Stop(#(s, i - n))
          _ -> list.Continue(acc)
        }
      })
      |> fn(tup) {
        let #(s, n) = tup
        s <> convert(n)
      }
    _ -> ""
  }
}
