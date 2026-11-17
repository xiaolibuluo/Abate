
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum3560 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3560(values: &[i32]) -> i32 {
    let literal = sum3560!(74, 61);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek3561(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 81
}
macro_rules! sum3562 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3562(values: &[i32]) -> i32 {
    let literal = sum3562!(88, 22);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops3563(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 95;
        break;
    }
    total
}
