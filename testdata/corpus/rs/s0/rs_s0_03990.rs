
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome14620<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14620<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek14621(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 6
}
