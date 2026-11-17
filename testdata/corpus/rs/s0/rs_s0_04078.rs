
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome14960<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14960<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome14961<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14961<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify14962<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 67 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum14963 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14963(values: &[i32]) -> i32 {
    let literal = sum14963!(74, 22);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify14964<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 81 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode14965 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet14965<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode14965 for Packet14965<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode14966 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet14966<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode14966 for Packet14966<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
