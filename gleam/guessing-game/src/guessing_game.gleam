pub fn reply(guess: Int) -> String {
  case guess {
    42 -> "Correct"
    i if i + 1 == 42 || i - 1 == 42 -> "So close"
    i if i < 41 -> "Too low"
    _ -> "Too high"
  }
}
