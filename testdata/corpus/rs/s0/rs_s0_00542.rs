
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode1950 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1950<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1950 for Packet1950<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
