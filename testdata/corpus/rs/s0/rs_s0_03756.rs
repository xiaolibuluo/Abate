
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum13780 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13780(values: &[i32]) -> i32 {
    let literal = sum13780!(71, 17);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify13781<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 78 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode13782 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13782<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13782 for Packet13782<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode13783 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13783<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13783 for Packet13783<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
