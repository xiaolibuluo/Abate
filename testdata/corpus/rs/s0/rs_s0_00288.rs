
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode990 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet990<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode990 for Packet990<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
