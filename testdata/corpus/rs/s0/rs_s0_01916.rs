
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome7090<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7090<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify7091<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 14 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum7092 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7092(values: &[i32]) -> i32 {
    let literal = sum7092!(21, 53);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum7093 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7093(values: &[i32]) -> i32 {
    let literal = sum7093!(28, 64);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum7094 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7094(values: &[i32]) -> i32 {
    let literal = sum7094!(35, 14);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek7095(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 42
}
