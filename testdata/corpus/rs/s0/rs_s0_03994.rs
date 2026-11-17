
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome14650<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14650<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome14651<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14651<T> {
    pub value: T,
    pub tag: u32,
}
