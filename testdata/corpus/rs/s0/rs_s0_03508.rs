
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum12910 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12910(values: &[i32]) -> i32 {
    let literal = sum12910!(7, 18);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum12911 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12911(values: &[i32]) -> i32 {
    let literal = sum12911!(14, 29);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode12912 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet12912<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode12912 for Packet12912<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode12913 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet12913<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode12913 for Packet12913<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek12914(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 35
}
