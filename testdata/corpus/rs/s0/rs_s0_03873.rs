
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse14200(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 35)
}
pub fn peek14201(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 42
}
#[derive(Debug, Clone)]
pub enum Outcome14202<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14202<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops14203(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 56;
        break;
    }
    total
}
