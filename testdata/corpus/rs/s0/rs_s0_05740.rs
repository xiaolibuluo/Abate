
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek20780(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 53
}
pub trait Encode20781 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20781<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20781 for Packet20781<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode20782 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20782<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20782 for Packet20782<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify20783<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 74 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode20784 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20784<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20784 for Packet20784<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum20785 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20785(values: &[i32]) -> i32 {
    let literal = sum20785!(88, 31);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify20786<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 95 => name,
        Some(_) => "other",
        None => "empty",
    }
}
