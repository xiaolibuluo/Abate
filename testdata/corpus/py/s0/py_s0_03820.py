
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1761_0(value: int) -> int:
    return value + 6


def decorate1761_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1761_0
def doubled1761_0(value: int) -> int:
    return value + value
class Slot1761_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{13}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1761_2(value: int) -> int:
    return value + 20


def decorate1761_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1761_2
def doubled1761_2(value: int) -> int:
    return value + value
