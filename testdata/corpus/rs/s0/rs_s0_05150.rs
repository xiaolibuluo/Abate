
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek18590(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
macro_rules! sum18591 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18591(values: &[i32]) -> i32 {
    let literal = sum18591!(26, 40);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse18592(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 33)
}
pub fn classify18593<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 40 => name,
        Some(_) => "other",
        None => "empty",
    }
}
