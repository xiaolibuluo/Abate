
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode2090 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2090<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2090 for Packet2090<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode2091 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2091<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2091 for Packet2091<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek2092(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 20
}
pub trait Encode2093 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2093<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2093 for Packet2093<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
