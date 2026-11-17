
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome9720<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9720<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode9721 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet9721<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode9721 for Packet9721<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome9722<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9722<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome9723<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9723<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode9724 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet9724<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode9724 for Packet9724<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode9725 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet9725<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode9725 for Packet9725<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify9726<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 73 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse9727(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 80)
}
pub fn classify9728<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 87 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek9729(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 94
}
