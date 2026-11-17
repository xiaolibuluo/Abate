
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode12190 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet12190<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode12190 for Packet12190<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
