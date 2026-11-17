
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome9100<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9100<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome9101<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9101<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek9102(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 15
}
pub fn peek9103(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 22
}
pub trait Encode9104 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet9104<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode9104 for Packet9104<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome9105<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9105<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome9106<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9106<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify9107<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 50 => name,
        Some(_) => "other",
        None => "empty",
    }
}
