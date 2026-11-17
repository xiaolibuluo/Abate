
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify2010<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 96 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum2011 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2011(values: &[i32]) -> i32 {
    let literal = sum2011!(6, 24);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
