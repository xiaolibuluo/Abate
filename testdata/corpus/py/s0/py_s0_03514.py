
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1621_0(values):
    for item in values:
        if item > 29:
            yield item * 10


def collect1621_0(values):
    return list(expand1621_0(values))
def traced1621_1(value: int) -> int:
    return value + 36


def decorate1621_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1621_1
def doubled1621_1(value: int) -> int:
    return value + value
class Slot1621_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{43}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1621_3(values):
    for item in values:
        if item > 50:
            yield item * 43


def collect1621_3(values):
    return list(expand1621_3(values))
