
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum9270 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9270(values: &[i32]) -> i32 {
    let literal = sum9270!(28, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum9271 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9271(values: &[i32]) -> i32 {
    let literal = sum9271!(35, 43);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
