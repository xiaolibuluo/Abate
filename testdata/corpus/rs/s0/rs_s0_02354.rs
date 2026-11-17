
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops8850(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 64;
        break;
    }
    total
}
macro_rules! sum8851 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8851(values: &[i32]) -> i32 {
    let literal = sum8851!(71, 54);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
