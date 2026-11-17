
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome9300<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9300<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome9301<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9301<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome9302<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9302<T> {
    pub value: T,
    pub tag: u32,
}
