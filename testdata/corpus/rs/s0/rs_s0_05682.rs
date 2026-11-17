
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify20560<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 58 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek20561(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 65
}
#[derive(Debug, Clone)]
pub enum Outcome20562<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20562<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek20563(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 79
}
