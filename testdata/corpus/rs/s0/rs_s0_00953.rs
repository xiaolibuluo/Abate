
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome3490<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3490<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome3491<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3491<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek3492(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 94
}
