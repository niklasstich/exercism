use std::collections::HashSet;

//"iterative" approach
pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[& 'a str]) -> HashSet<&'a str> {
    //get lowercase and sorted version of original word once and reuse it
    let word_lowercase = word.to_lowercase();
    let mut word_charvec: Vec<char> = word_lowercase.chars().collect();
    word_charvec.sort_unstable();
    let word_sorted: String = word_charvec.iter().collect();
    //construct hashset for our results
    let mut retval = HashSet::new();
    for anagram_candidate in possible_anagrams.iter() {
        if check_anagram(word_lowercase.as_str(), word_sorted.as_str(), anagram_candidate) {
            retval.insert(*anagram_candidate);
        }
    }

    return retval;
}

fn check_anagram(word_lc: &str, word_sorted: &str, candidate: &str) -> bool {
    //anagrams must be the same length
    if word_lc.len() != candidate.len() {
        return false;
    }
    let candidate_lc = candidate.to_lowercase();

    //if the char vectors are identical here before sorting, that means the supplied words are identical
    //it is therefore not an anagram as per exercise description
    if word_lc == candidate_lc {
        return false;
    }

    let mut candidate_chars: Vec<char> = candidate_lc.chars().collect();
    candidate_chars.sort_unstable();
    //collect char vector back into string
    let candidate_sorted: String = candidate_chars.iter().collect();

    //if the sorted strings are equal, its an anagram
    word_sorted == candidate_sorted
}
