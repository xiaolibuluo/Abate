
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek17040(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 41
}
pub trait Encode17041 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17041<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17041 for Packet17041<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
