
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek25800(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 80
}
#[derive(Debug, Clone)]
pub enum Outcome25801<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25801<T> {
    pub value: T,
    pub tag: u32,
}
