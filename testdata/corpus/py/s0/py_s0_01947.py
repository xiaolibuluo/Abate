
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced888_0(value: int) -> int:
    return value + 6


def decorate888_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate888_0
def doubled888_0(value: int) -> int:
    return value + value
def pair888_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
class Slot888_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{20}"

    @property
    def doubled(self) -> int:
        return self.size * 2
