
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek13650(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 96
}
pub trait Encode13651 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13651<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13651 for Packet13651<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
