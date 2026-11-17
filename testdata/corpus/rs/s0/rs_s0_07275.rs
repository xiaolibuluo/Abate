
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify26210<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 31 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse26211(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 38)
}
pub fn classify26212<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 45 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse26213(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 52)
}
#[derive(Debug, Clone)]
pub enum Outcome26214<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record26214<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify26215<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 66 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum26216 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold26216(values: &[i32]) -> i32 {
    let literal = sum26216!(73, 61);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome26217<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record26217<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify26218<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 87 => name,
        Some(_) => "other",
        None => "empty",
    }
}
