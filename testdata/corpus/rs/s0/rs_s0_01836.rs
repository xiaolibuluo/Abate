
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek6790(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 5
}
macro_rules! sum6791 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold6791(values: &[i32]) -> i32 {
    let literal = sum6791!(12, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse6792(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 19)
}
pub fn parse6793(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 26)
}
pub fn parse6794(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 33)
}
pub trait Encode6795 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6795<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6795 for Packet6795<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn loops6796(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 47;
        break;
    }
    total
}
pub fn loops6797(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 54;
        break;
    }
    total
}
pub fn peek6798(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 61
}
macro_rules! sum6799 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold6799(values: &[i32]) -> i32 {
    let literal = sum6799!(68, 42);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
