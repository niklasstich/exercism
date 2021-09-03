pub fn build_proverb(list: &[&str]) -> String {
    let mut retval = String::new();
    if list.is_empty() {
        return retval
    }
    let mut iter = list.iter().peekable();
    loop {
        let w1 = iter.next().unwrap();
        let w2 = iter.peek();
        if w2.is_none() {
            break;
        }
        retval.push_str(&format!("For want of a {} the {} was lost.\n", w1, w2.unwrap()));
    }
    retval.push_str(&format!("And all for the want of a {}.", list.first().unwrap()));

    retval
}
