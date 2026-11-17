
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum5400 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5400(values: &[i32]) -> i32 {
    let literal = sum5400!(41, 68);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify5401<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 48 => name,
        Some(_) => "other",
        None => "empty",
    }
}
