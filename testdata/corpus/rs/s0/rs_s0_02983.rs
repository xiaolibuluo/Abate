
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek10970(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 7
}
#[derive(Debug, Clone)]
pub enum Outcome10971<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10971<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify10972<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 21 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome10973<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10973<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek10974(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 35
}
pub fn peek10975(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 42
}
pub fn peek10976(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 49
}
pub fn peek10977(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 56
}
