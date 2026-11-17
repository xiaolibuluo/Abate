
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify16370<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 43 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse16371(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 50)
}
#[derive(Debug, Clone)]
pub enum Outcome16372<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16372<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify16373<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 64 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify16374<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 71 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum16375 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16375(values: &[i32]) -> i32 {
    let literal = sum16375!(78, 55);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum16376 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16376(values: &[i32]) -> i32 {
    let literal = sum16376!(85, 66);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
