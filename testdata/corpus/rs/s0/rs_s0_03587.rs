
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome13190<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13190<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse13191(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 87)
}
pub fn loops13192(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 94;
        break;
    }
    total
}
