
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops3480(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 67;
        break;
    }
    total
}
pub fn peek3481(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 74
}
macro_rules! sum3482 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3482(values: &[i32]) -> i32 {
    let literal = sum3482!(81, 27);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
