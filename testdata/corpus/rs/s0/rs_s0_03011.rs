
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek11070(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 40
}
#[derive(Debug, Clone)]
pub enum Outcome11071<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11071<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode11072 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet11072<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode11072 for Packet11072<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify11073<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 61 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome11074<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11074<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum11075 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold11075(values: &[i32]) -> i32 {
    let literal = sum11075!(75, 66);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome11076<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11076<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek11077(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 89
}
