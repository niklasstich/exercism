pub fn raindrops(n: u32) -> String {
    let mut retval = String::new();
    if n % 3 == 0 {
        retval.push_str("Pling");
    }
    if n % 5 == 0 {
        retval.push_str("Plang");
    }
    if n % 7 == 0 {
        retval.push_str("Plong");
    }
    if !retval.is_empty() {
        return retval
    }
    n.to_string()
}
