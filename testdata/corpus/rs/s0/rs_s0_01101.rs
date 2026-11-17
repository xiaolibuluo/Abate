
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek4000(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 64
}
pub trait Encode4001 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4001<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4001 for Packet4001<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode4002 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4002<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4002 for Packet4002<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
