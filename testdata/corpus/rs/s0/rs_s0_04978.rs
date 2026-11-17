
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek18020(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 54
}
pub trait Encode18021 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet18021<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode18021 for Packet18021<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
