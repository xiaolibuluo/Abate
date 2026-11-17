
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode15630 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15630<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15630 for Packet15630<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome15631<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15631<T> {
    pub value: T,
    pub tag: u32,
}
