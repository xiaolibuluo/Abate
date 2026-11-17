
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek2570(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 48
}
pub trait Encode2571 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2571<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2571 for Packet2571<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode2572 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2572<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2572 for Packet2572<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
