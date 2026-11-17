
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome21810<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21810<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify21811<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 41 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek21812(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 48
}
pub fn classify21813<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 55 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum21814 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21814(values: &[i32]) -> i32 {
    let literal = sum21814!(62, 9);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify21815<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 69 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum21816 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21816(values: &[i32]) -> i32 {
    let literal = sum21816!(76, 31);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome21817<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21817<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek21818(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 90
}
