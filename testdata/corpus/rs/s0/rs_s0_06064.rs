
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode21920 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21920<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21920 for Packet21920<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
