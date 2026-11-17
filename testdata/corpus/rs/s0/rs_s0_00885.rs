
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum3250 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3250(values: &[i32]) -> i32 {
    let literal = sum3250!(59, 27);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify3251<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 66 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome3252<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3252<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome3253<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3253<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek3254(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 87
}
macro_rules! sum3255 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3255(values: &[i32]) -> i32 {
    let literal = sum3255!(94, 21);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum3256 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3256(values: &[i32]) -> i32 {
    let literal = sum3256!(101, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
