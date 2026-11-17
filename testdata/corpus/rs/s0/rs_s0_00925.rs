
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum3410 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3410(values: &[i32]) -> i32 {
    let literal = sum3410!(73, 17);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
