
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome22000<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22000<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome22001<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22001<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek22002(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 101
}
pub fn classify22003<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 11 => name,
        Some(_) => "other",
        None => "empty",
    }
}
