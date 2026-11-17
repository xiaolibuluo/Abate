
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode20070 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20070<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20070 for Packet20070<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode20071 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20071<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20071 for Packet20071<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek20072(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 17
}
