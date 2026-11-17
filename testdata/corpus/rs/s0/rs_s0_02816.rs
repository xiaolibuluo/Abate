
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops10400(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 42;
        break;
    }
    total
}
pub fn parse10401(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 49)
}
pub trait Encode10402 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10402<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10402 for Packet10402<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek10403(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 63
}
macro_rules! sum10404 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10404(values: &[i32]) -> i32 {
    let literal = sum10404!(70, 13);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
