use std::collections::HashSet;

//"iterative" approach
pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[& 'a str]) -> HashSet<&'a str> {
    let mut hs = HashSet::new();
    for anagram_candidate in possible_anagrams.iter() {
        if check_anagram(word, anagram_candidate) {
            hs.insert(*anagram_candidate);
        }
    }

    return hs;
}

fn check_anagram(w1: &str, w2: &str) -> bool {
    //anagrams must be the same length
    if w1.len() != w2.len() {
        return false;
    }

    //get char vectors of both words
    let mut w1_chars: Vec<char> = string_lowercase(w1).chars().collect();
    let mut w2_chars: Vec<char> = string_lowercase(w2).chars().collect();
    //if they are the same before sorting, they are already sorted, which means they are the same
    if w1_chars == w2_chars {
        return false;
    }
    //sort the vectors alphabetically
    w1_chars.sort_by(|a, b| b.cmp(a));
    w2_chars.sort_by(|a, b| b.cmp(a));

    //if the vectors are equal, its an anagram
    return w1_chars==w2_chars;
}


//"borrowed" functions below from https://exercism.io/tracks/rust/exercises/anagram/solutions/3cd55872b53048ecb3756daca3a877ae
//returns the lowercase version of a string
fn string_lowercase(str: &str) -> String {
    //collect chars from string into slice, map every char to its lowercase version, collect back into string
    str.chars().map(|c| char_lowercase(c)).collect()
}

//returns the lowercase version of a char
fn char_lowercase(c: char) -> char {
    //to_lowercase returns an iterator, so get next and unwrap value from it
    //we ignore when chars have a lowercase representation of more than one char
    c.to_lowercase().next().unwrap()
}
