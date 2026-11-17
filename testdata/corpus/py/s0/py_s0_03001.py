
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1374_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{19}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1374_1(value: int) -> int:
    return value + 26


def decorate1374_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1374_1
def doubled1374_1(value: int) -> int:
    return value + value
def expand1374_2(values):
    for item in values:
        if item > 33:
            yield item * 11


def collect1374_2(values):
    return list(expand1374_2(values))
