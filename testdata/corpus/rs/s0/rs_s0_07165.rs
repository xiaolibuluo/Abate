
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum25850 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold25850(values: &[i32]) -> i32 {
    let literal = sum25850!(48, 48);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify25851<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 55 => name,
        Some(_) => "other",
        None => "empty",
    }
}
