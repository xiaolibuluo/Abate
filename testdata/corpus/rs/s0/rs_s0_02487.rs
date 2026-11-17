
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify9330<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 9 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse9331(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 16)
}
macro_rules! sum9332 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9332(values: &[i32]) -> i32 {
    let literal = sum9332!(23, 35);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
