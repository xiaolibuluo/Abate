
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum24550 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24550(values: &[i32]) -> i32 {
    let literal = sum24550!(7, 53);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome24551<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24551<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify24552<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 21 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify24553<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 28 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome24554<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24554<T> {
    pub value: T,
    pub tag: u32,
}
