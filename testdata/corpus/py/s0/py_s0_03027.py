
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1388_0(value: int) -> int:
    return value + 7


def decorate1388_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1388_0
def doubled1388_0(value: int) -> int:
    return value + value
def traced1388_1(value: int) -> int:
    return value + 14


def decorate1388_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1388_1
def doubled1388_1(value: int) -> int:
    return value + value
def pair1388_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
def traced1388_3(value: int) -> int:
    return value + 28


def decorate1388_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1388_3
def doubled1388_3(value: int) -> int:
    return value + value
class Slot1388_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{35}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1388_5(values):
    for item in values:
        if item > 42:
            yield item * 20


def collect1388_5(values):
    return list(expand1388_5(values))
