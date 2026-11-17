
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek23900(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 35
}
pub trait Encode23901 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23901<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23901 for Packet23901<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
