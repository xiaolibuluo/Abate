
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode24990 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet24990<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode24990 for Packet24990<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode24991 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet24991<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode24991 for Packet24991<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode24992 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet24992<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode24992 for Packet24992<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
