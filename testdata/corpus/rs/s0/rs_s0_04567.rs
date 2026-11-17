
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome16580<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16580<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek16581(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 32
}
