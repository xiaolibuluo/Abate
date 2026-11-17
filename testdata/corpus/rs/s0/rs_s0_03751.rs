
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum13760 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13760(values: &[i32]) -> i32 {
    let literal = sum13760!(45, 64);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome13761<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13761<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek13762(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 59
}
pub fn classify13763<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 66 => name,
        Some(_) => "other",
        None => "empty",
    }
}
