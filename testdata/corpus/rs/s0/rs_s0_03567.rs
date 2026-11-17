
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse13130(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 99)
}
pub trait Encode13131 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13131<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13131 for Packet13131<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek13132(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 16
}
pub trait Encode13133 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13133<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13133 for Packet13133<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum13134 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13134(values: &[i32]) -> i32 {
    let literal = sum13134!(30, 33);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum13135 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13135(values: &[i32]) -> i32 {
    let literal = sum13135!(37, 44);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome13136<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13136<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome13137<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13137<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse13138(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 58)
}
macro_rules! sum13139 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13139(values: &[i32]) -> i32 {
    let literal = sum13139!(65, 27);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
