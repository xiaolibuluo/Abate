
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify17440<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 76 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum17441 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17441(values: &[i32]) -> i32 {
    let literal = sum17441!(83, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify17442<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 90 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek17443(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 97
}
pub fn peek17444(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 7
}
pub fn classify17445<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 14 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum17446 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17446(values: &[i32]) -> i32 {
    let literal = sum17446!(21, 22);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify17447<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 28 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse17448(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 35)
}
