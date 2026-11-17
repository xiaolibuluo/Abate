
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek13840(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 52
}
pub trait Encode13841 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13841<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13841 for Packet13841<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum13842 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13842(values: &[i32]) -> i32 {
    let literal = sum13842!(66, 20);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode13843 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13843<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13843 for Packet13843<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum13844 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13844(values: &[i32]) -> i32 {
    let literal = sum13844!(80, 42);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
