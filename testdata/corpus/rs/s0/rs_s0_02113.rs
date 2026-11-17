
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode7840 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet7840<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode7840 for Packet7840<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek7841(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
