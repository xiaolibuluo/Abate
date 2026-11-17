
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify21580<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 26 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse21581(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 33)
}
pub trait Encode21582 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21582<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21582 for Packet21582<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum21583 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21583(values: &[i32]) -> i32 {
    let literal = sum21583!(47, 20);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
