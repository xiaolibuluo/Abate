
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode2120 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2120<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2120 for Packet2120<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn loops2121(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 52;
        break;
    }
    total
}
macro_rules! sum2122 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2122(values: &[i32]) -> i32 {
    let literal = sum2122!(59, 51);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum2123 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2123(values: &[i32]) -> i32 {
    let literal = sum2123!(66, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
