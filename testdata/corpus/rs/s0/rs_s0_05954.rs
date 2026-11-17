
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode21540 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21540<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21540 for Packet21540<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek21541(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 78
}
macro_rules! sum21542 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21542(values: &[i32]) -> i32 {
    let literal = sum21542!(85, 42);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum21543 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21543(values: &[i32]) -> i32 {
    let literal = sum21543!(92, 53);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
