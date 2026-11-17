
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify10940<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 65 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse10941(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 72)
}
pub fn peek10942(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 79
}
#[derive(Debug, Clone)]
pub enum Outcome10943<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10943<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum10944 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10944(values: &[i32]) -> i32 {
    let literal = sum10944!(93, 25);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum10945 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10945(values: &[i32]) -> i32 {
    let literal = sum10945!(100, 36);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
