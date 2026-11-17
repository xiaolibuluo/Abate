
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek23600(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 33
}
pub fn peek23601(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 40
}
#[derive(Debug, Clone)]
pub enum Outcome23602<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23602<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek23603(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 54
}
#[derive(Debug, Clone)]
pub enum Outcome23604<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23604<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify23605<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 68 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek23606(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 75
}
