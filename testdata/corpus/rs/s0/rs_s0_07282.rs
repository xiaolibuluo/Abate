
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops26270(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 12;
        break;
    }
    total
}
pub fn parse26271(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 19)
}
pub fn loops26272(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 26;
        break;
    }
    total
}
pub fn peek26273(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 33
}
pub fn classify26274<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 40 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse26275(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 47)
}
pub trait Encode26276 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet26276<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode26276 for Packet26276<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum26277 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold26277(values: &[i32]) -> i32 {
    let literal = sum26277!(61, 53);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek26278(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 68
}
pub trait Encode26279 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet26279<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode26279 for Packet26279<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
