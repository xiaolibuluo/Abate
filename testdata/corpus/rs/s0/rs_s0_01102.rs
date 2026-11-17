
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum4010 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4010(values: &[i32]) -> i32 {
    let literal = sum4010!(77, 10);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum4011 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4011(values: &[i32]) -> i32 {
    let literal = sum4011!(84, 21);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify4012<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 91 => name,
        Some(_) => "other",
        None => "empty",
    }
}
