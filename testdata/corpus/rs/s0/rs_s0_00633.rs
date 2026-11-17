
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek2230(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 91
}
pub fn parse2231(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 98)
}
pub fn classify2232<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 8 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode2233 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2233<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2233 for Packet2233<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum2234 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2234(values: &[i32]) -> i32 {
    let literal = sum2234!(22, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek2235(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 29
}
pub trait Encode2236 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2236<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2236 for Packet2236<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
