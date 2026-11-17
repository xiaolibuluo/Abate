
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode25410 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet25410<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode25410 for Packet25410<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode25411 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet25411<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode25411 for Packet25411<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek25412(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 72
}
