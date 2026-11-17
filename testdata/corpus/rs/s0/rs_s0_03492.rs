
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome12860<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12860<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum12861 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12861(values: &[i32]) -> i32 {
    let literal = sum12861!(46, 55);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode12862 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet12862<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode12862 for Packet12862<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek12863(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 60
}
macro_rules! sum12864 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12864(values: &[i32]) -> i32 {
    let literal = sum12864!(67, 27);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
