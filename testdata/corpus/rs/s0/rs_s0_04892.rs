
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek17730(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 65
}
pub fn parse17731(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 72)
}
pub fn peek17732(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 79
}
macro_rules! sum17733 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17733(values: &[i32]) -> i32 {
    let literal = sum17733!(86, 9);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops17734(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 93;
        break;
    }
    total
}
pub fn peek17735(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 100
}
