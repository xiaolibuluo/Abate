
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek18040(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 80
}
pub trait Encode18041 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet18041<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode18041 for Packet18041<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
