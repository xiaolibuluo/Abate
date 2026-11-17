
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode16090 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet16090<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode16090 for Packet16090<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek16091(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 74
}
