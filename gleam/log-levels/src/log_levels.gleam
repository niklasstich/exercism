import gleam/string

pub fn message(log_line: String) -> String {
  case string.split(log_line, ": ") {
    [_, s] -> s |> string.trim
    _ -> ""
  }
}

pub fn log_level(log_line: String) -> String {
  case string.split(log_line, ": ") {
    [s, _] ->
      s |> string.drop_left(1) |> string.drop_right(1) |> string.lowercase
    _ -> ""
  }
}

pub fn reformat(log_line: String) -> String {
  let msg = message(log_line)
  let lvl = log_level(log_line)
  msg <> " (" <> lvl <> ")"
}
