
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse15860(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 59)
}
pub fn parse15861(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 66)
}
pub fn classify15862<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 73 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode15863 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15863<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15863 for Packet15863<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek15864(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 87
}
macro_rules! sum15865 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15865(values: &[i32]) -> i32 {
    let literal = sum15865!(94, 64);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum15866 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15866(values: &[i32]) -> i32 {
    let literal = sum15866!(101, 14);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
