
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode9740 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet9740<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode9740 for Packet9740<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
