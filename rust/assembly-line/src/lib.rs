// This stub file contains items which aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

pub fn production_rate_per_hour(speed: u8) -> f64 {
    let prod_factor = match speed {
        1..=4 => 1.0,
        5..=8 => 0.9,
        9 | 10 => 0.77,
        _ => 0.0,
    };
    let raw_prod = 221 * speed as u32;
    prod_factor * (raw_prod as f64)
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
}
