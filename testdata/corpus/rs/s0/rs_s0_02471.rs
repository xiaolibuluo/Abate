
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek9290(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 54
}
pub trait Encode9291 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet9291<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode9291 for Packet9291<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
