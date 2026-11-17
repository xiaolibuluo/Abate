
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify13070<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 21 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum13071 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13071(values: &[i32]) -> i32 {
    let literal = sum13071!(28, 19);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify13072<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 35 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum13073 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13073(values: &[i32]) -> i32 {
    let literal = sum13073!(42, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify13074<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 49 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode13075 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13075<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13075 for Packet13075<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify13076<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 63 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek13077(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 70
}
