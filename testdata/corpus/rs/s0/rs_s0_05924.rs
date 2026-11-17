
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome21460<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21460<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek21461(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 71
}
#[derive(Debug, Clone)]
pub enum Outcome21462<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21462<T> {
    pub value: T,
    pub tag: u32,
}
