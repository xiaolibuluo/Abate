
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify24410<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 19 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek24411(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 26
}
#[derive(Debug, Clone)]
pub enum Outcome24412<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24412<T> {
    pub value: T,
    pub tag: u32,
}
