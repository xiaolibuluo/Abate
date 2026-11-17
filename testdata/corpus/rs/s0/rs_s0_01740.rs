
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome6410<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record6410<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify6411<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 100 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode6412 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6412<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6412 for Packet6412<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
