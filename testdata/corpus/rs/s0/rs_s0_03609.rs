
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode13270 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13270<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13270 for Packet13270<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek13271(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 94
}
pub trait Encode13272 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13272<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13272 for Packet13272<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
