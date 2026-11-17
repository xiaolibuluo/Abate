
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode17810 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17810<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17810 for Packet17810<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek17811(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 79
}
pub fn classify17812<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 86 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum17813 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17813(values: &[i32]) -> i32 {
    let literal = sum17813!(93, 65);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum17814 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17814(values: &[i32]) -> i32 {
    let literal = sum17814!(100, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek17815(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 10
}
pub fn peek17816(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 17
}
