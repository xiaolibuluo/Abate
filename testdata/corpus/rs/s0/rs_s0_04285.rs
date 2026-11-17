
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops15710(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 58;
        break;
    }
    total
}
pub fn peek15711(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 65
}
pub trait Encode15712 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15712<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15712 for Packet15712<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum15713 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15713(values: &[i32]) -> i32 {
    let literal = sum15713!(79, 59);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
