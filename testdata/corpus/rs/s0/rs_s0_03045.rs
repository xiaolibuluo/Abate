
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode11190 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet11190<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode11190 for Packet11190<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
