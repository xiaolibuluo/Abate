
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode6090 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6090<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6090 for Packet6090<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek6091(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 72
}
pub trait Encode6092 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6092<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6092 for Packet6092<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode6093 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6093<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6093 for Packet6093<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode6094 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6094<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6094 for Packet6094<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
