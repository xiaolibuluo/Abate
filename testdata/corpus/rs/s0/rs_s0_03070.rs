
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode11250 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet11250<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode11250 for Packet11250<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome11251<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11251<T> {
    pub value: T,
    pub tag: u32,
}
