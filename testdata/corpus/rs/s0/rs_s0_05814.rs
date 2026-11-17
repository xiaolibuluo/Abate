
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode21050 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21050<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21050 for Packet21050<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode21051 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21051<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21051 for Packet21051<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode21052 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21052<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21052 for Packet21052<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
