
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome12670<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12670<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome12671<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12671<T> {
    pub value: T,
    pub tag: u32,
}
