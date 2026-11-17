
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse24520(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 65)
}
pub fn peek24521(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 72
}
pub fn loops24522(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 79;
        break;
    }
    total
}
macro_rules! sum24523 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24523(values: &[i32]) -> i32 {
    let literal = sum24523!(86, 65);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum24524 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24524(values: &[i32]) -> i32 {
    let literal = sum24524!(93, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
