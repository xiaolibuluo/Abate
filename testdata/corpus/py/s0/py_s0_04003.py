
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1858_0(values):
    for item in values:
        if item > 6:
            yield item * 22


def collect1858_0(values):
    return list(expand1858_0(values))
class Slot1858_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{13}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1858_2(value: int) -> int:
    return value + 20


def decorate1858_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1858_2
def doubled1858_2(value: int) -> int:
    return value + value
