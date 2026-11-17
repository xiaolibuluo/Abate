
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome18170<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18170<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome18171<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18171<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum18172 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18172(values: &[i32]) -> i32 {
    let literal = sum18172!(69, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify18173<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 76 => name,
        Some(_) => "other",
        None => "empty",
    }
}
