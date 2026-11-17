
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode7000 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet7000<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode7000 for Packet7000<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek7001(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 91
}
pub trait Encode7002 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet7002<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode7002 for Packet7002<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek7003(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 8
}
macro_rules! sum7004 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7004(values: &[i32]) -> i32 {
    let literal = sum7004!(15, 12);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
