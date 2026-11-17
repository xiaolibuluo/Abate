
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify20540<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 32 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify20541<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 39 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum20542 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20542(values: &[i32]) -> i32 {
    let literal = sum20542!(46, 13);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
