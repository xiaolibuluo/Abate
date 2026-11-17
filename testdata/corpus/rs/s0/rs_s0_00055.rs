
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops160(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 19;
        break;
    }
    total
}
macro_rules! sum161 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold161(values: &[i32]) -> i32 {
    let literal = sum161!(26, 10);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum162 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold162(values: &[i32]) -> i32 {
    let literal = sum162!(33, 21);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum163 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold163(values: &[i32]) -> i32 {
    let literal = sum163!(40, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
