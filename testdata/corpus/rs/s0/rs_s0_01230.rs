
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify4490<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 22 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify4491<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 29 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum4492 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4492(values: &[i32]) -> i32 {
    let literal = sum4492!(36, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
