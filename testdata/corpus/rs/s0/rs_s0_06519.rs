
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek23480(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 71
}
pub trait Encode23481 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23481<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23481 for Packet23481<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum23482 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23482(values: &[i32]) -> i32 {
    let literal = sum23482!(85, 58);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
