
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode25810 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet25810<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode25810 for Packet25810<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
