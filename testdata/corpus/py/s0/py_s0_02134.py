
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced977_0(value: int) -> int:
    return value + 96


def decorate977_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate977_0
def doubled977_0(value: int) -> int:
    return value + value
class Slot977_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{6}"

    @property
    def doubled(self) -> int:
        return self.size * 2
