
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify10380<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 16 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode10381 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10381<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10381 for Packet10381<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum10382 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10382(values: &[i32]) -> i32 {
    let literal = sum10382!(30, 38);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse10383(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 37)
}
