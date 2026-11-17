
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek6350(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 15
}
pub trait Encode6351 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6351<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6351 for Packet6351<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
