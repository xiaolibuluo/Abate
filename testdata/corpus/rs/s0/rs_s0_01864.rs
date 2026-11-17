
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek6860(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 96
}
pub trait Encode6861 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6861<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6861 for Packet6861<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
