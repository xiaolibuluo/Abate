
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15460<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15460<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops15461(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 31;
        break;
    }
    total
}
pub fn peek15462(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 38
}
macro_rules! sum15463 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15463(values: &[i32]) -> i32 {
    let literal = sum15463!(45, 67);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode15464 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15464<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15464 for Packet15464<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum15465 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15465(values: &[i32]) -> i32 {
    let literal = sum15465!(59, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify15466<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 66 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek15467(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 73
}
