
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode13210 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13210<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13210 for Packet13210<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode13211 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13211<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13211 for Packet13211<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
