
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum14280 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14280(values: &[i32]) -> i32 {
    let literal = sum14280!(42, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify14281<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 49 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse14282(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 56)
}
pub fn peek14283(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 63
}
pub trait Encode14284 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet14284<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode14284 for Packet14284<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
