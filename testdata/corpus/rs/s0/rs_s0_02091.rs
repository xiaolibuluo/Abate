
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek7780(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 31
}
pub fn peek7781(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 38
}
pub trait Encode7782 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet7782<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode7782 for Packet7782<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek7783(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 52
}
pub fn peek7784(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 59
}
pub fn classify7785<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 66 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek7786(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 73
}
pub fn peek7787(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 80
}
pub fn classify7788<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 87 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode7789 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet7789<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode7789 for Packet7789<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse77810(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 101)
}
macro_rules! sum77811 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold77811(values: &[i32]) -> i32 {
    let literal = sum77811!(11, 25);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
