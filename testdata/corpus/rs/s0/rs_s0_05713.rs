
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome20680<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20680<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum20681 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20681(values: &[i32]) -> i32 {
    let literal = sum20681!(27, 39);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum20682 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20682(values: &[i32]) -> i32 {
    let literal = sum20682!(34, 50);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify20683<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 41 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum20684 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20684(values: &[i32]) -> i32 {
    let literal = sum20684!(48, 11);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
