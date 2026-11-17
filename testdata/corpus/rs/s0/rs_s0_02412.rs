
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum9070 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9070(values: &[i32]) -> i32 {
    let literal = sum9070!(59, 14);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome9071<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9071<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify9072<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 73 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek9073(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 80
}
#[derive(Debug, Clone)]
pub enum Outcome9074<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9074<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify9075<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 94 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum9076 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9076(values: &[i32]) -> i32 {
    let literal = sum9076!(101, 19);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
