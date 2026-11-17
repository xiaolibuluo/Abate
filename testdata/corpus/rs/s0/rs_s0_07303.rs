
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode26340 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet26340<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode26340 for Packet26340<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode26341 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet26341<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode26341 for Packet26341<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
