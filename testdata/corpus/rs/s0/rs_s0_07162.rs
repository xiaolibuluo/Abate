
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome25830<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25830<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome25831<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25831<T> {
    pub value: T,
    pub tag: u32,
}
