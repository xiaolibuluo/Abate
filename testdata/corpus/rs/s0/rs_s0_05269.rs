
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome18980<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18980<T> {
    pub value: T,
    pub tag: u32,
}
