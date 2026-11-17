
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek24830(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 80
}
#[derive(Debug, Clone)]
pub enum Outcome24831<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24831<T> {
    pub value: T,
    pub tag: u32,
}
