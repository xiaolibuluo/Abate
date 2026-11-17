
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome18510<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18510<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome18511<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18511<T> {
    pub value: T,
    pub tag: u32,
}
