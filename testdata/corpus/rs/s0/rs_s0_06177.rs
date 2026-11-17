
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode22380 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet22380<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode22380 for Packet22380<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
