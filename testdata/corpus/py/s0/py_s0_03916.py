
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1814_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{16}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1814_1(value: int) -> int:
    return value + 23


def decorate1814_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1814_1
def doubled1814_1(value: int) -> int:
    return value + value
def expand1814_2(values):
    for item in values:
        if item > 30:
            yield item * 41


def collect1814_2(values):
    return list(expand1814_2(values))
def project1814_3(values):
    select = lambda item: item + 37
    return [select(item) for item in values if item != 0]
