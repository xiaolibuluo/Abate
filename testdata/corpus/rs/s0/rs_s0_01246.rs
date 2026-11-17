
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops4560(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 16;
        break;
    }
    total
}
pub fn peek4561(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 23
}
pub fn parse4562(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 30)
}
pub fn loops4563(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 37;
        break;
    }
    total
}
macro_rules! sum4564 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4564(values: &[i32]) -> i32 {
    let literal = sum4564!(44, 12);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
