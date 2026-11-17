
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode15820 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15820<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15820 for Packet15820<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify15821<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 14 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify15822<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 21 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode15823 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15823<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15823 for Packet15823<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum15824 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15824(values: &[i32]) -> i32 {
    let literal = sum15824!(35, 25);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
