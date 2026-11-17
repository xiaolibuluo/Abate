
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek15000(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 8
}
pub fn parse15001(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 15)
}
pub fn classify15002<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 22 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify15003<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 29 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum15004 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15004(values: &[i32]) -> i32 {
    let literal = sum15004!(36, 61);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify15005<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 43 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek15006(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 50
}
pub fn classify15007<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 57 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek15008(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 64
}
pub fn parse15009(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 71)
}
