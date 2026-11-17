
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse6010(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 58)
}
pub fn peek6011(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 65
}
pub fn peek6012(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 72
}
pub fn classify6013<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 79 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse6014(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 86)
}
macro_rules! sum6015 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold6015(values: &[i32]) -> i32 {
    let literal = sum6015!(93, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
