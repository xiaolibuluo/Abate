
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1134_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{100}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1134_1(values):
    for item in values:
        if item > 10:
            yield item * 28


def collect1134_1(values):
    return list(expand1134_1(values))
def pair1134_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 17}
def traced1134_3(value: int) -> int:
    return value + 24


def decorate1134_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1134_3
def doubled1134_3(value: int) -> int:
    return value + value
