
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek23880(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 9
}
pub trait Encode23881 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23881<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23881 for Packet23881<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
