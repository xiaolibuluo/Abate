
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse13290(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 16)
}
pub fn classify13291<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 23 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse13292(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 30)
}
macro_rules! sum13293 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13293(values: &[i32]) -> i32 {
    let literal = sum13293!(37, 12);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek13294(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 44
}
