
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome16670<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16670<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum16671 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16671(values: &[i32]) -> i32 {
    let literal = sum16671!(52, 38);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify16672<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 59 => name,
        Some(_) => "other",
        None => "empty",
    }
}
