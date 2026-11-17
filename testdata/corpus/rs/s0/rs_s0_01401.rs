
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek5160(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 20
}
macro_rules! sum5161 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5161(values: &[i32]) -> i32 {
    let literal = sum5161!(27, 33);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode5162 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet5162<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode5162 for Packet5162<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum5163 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5163(values: &[i32]) -> i32 {
    let literal = sum5163!(41, 55);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum5164 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5164(values: &[i32]) -> i32 {
    let literal = sum5164!(48, 66);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome5165<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5165<T> {
    pub value: T,
    pub tag: u32,
}
