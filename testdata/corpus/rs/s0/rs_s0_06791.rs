
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome24510<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24510<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek24511(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 59
}
#[derive(Debug, Clone)]
pub enum Outcome24512<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24512<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome24513<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24513<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify24514<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 80 => name,
        Some(_) => "other",
        None => "empty",
    }
}
