
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek1850(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 82
}
pub fn classify1851<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 89 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek1852(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 96
}
pub trait Encode1853 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1853<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1853 for Packet1853<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum1854 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1854(values: &[i32]) -> i32 {
    let literal = sum1854!(13, 67);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify1855<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 20 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum1856 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1856(values: &[i32]) -> i32 {
    let literal = sum1856!(27, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum1857 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1857(values: &[i32]) -> i32 {
    let literal = sum1857!(34, 39);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
