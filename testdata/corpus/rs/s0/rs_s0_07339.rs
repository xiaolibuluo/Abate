
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops26480(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 91;
        break;
    }
    total
}
macro_rules! sum26481 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold26481(values: &[i32]) -> i32 {
    let literal = sum26481!(98, 12);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum26482 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold26482(values: &[i32]) -> i32 {
    let literal = sum26482!(8, 23);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops26483(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 15;
        break;
    }
    total
}
