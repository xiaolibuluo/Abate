
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek23190(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 82
}
pub fn peek23191(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 89
}
#[derive(Debug, Clone)]
pub enum Outcome23192<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23192<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode23193 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23193<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23193 for Packet23193<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum23194 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23194(values: &[i32]) -> i32 {
    let literal = sum23194!(13, 60);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum23195 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23195(values: &[i32]) -> i32 {
    let literal = sum23195!(20, 10);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode23196 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23196<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23196 for Packet23196<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
