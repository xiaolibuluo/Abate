
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek870(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 69
}
macro_rules! sum871 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold871(values: &[i32]) -> i32 {
    let literal = sum871!(76, 19);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify872<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 83 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify873<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 90 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum874 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold874(values: &[i32]) -> i32 {
    let literal = sum874!(97, 52);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse875(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 7)
}
pub fn classify876<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 14 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify877<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 21 => name,
        Some(_) => "other",
        None => "empty",
    }
}
