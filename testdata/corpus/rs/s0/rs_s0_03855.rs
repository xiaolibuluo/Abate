
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum14110 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14110(values: &[i32]) -> i32 {
    let literal = sum14110!(15, 65);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
