
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome17120<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17120<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode17121 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17121<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17121 for Packet17121<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify17122<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 62 => name,
        Some(_) => "other",
        None => "empty",
    }
}
