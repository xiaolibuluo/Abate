
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome19730<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record19730<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome19731<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record19731<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum19732 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold19732(values: &[i32]) -> i32 {
    let literal = sum19732!(60, 56);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum19733 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold19733(values: &[i32]) -> i32 {
    let literal = sum19733!(67, 67);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome19734<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record19734<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify19735<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 81 => name,
        Some(_) => "other",
        None => "empty",
    }
}
