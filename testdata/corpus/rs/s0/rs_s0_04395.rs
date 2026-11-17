
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode16100 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet16100<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode16100 for Packet16100<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode16101 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet16101<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode16101 for Packet16101<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
