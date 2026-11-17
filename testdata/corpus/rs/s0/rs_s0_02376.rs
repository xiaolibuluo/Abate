
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify8940<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 84 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome8941<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8941<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek8942(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 98
}
pub fn classify8943<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 8 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome8944<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8944<T> {
    pub value: T,
    pub tag: u32,
}
