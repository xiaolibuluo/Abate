
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15160<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15160<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome15161<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15161<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome15162<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15162<T> {
    pub value: T,
    pub tag: u32,
}
