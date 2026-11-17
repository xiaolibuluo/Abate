
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify12920<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 20 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify12921<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 27 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek12922(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 34
}
#[derive(Debug, Clone)]
pub enum Outcome12923<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12923<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify12924<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 48 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum12925 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12925(values: &[i32]) -> i32 {
    let literal = sum12925!(55, 19);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
