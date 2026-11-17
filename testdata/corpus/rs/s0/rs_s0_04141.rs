
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15220<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15220<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum15221 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15221(values: &[i32]) -> i32 {
    let literal = sum15221!(10, 60);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops15222(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 17;
        break;
    }
    total
}
