module TwoFer

let twoFer (input: string option): string = 
    (*Solution using pattern matching
    match input with 
    | Some(input) -> $"One for {input}, one for me."
    | None -> "One for you, one for me."
    *)

    //Solution using Option.defaultValue to unpack the option input
    ("you", input) ||> Option.defaultValue |> sprintf "One for %s, one for me."