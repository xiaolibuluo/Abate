
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek16080(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 54
}
pub trait Encode16081 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet16081<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode16081 for Packet16081<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
