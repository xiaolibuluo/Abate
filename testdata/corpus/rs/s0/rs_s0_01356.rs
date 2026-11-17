
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome4990<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4990<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode4991 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4991<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4991 for Packet4991<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn loops4992(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 7;
        break;
    }
    total
}
