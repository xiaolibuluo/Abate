
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1937_0(values):
    for item in values:
        if item > 63:
            yield item * 26


def collect1937_0(values):
    return list(expand1937_0(values))
def expand1937_1(values):
    for item in values:
        if item > 70:
            yield item * 37


def collect1937_1(values):
    return list(expand1937_1(values))
class Slot1937_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{77}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1937_3(value: int) -> int:
    return value + 84


def decorate1937_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1937_3
def doubled1937_3(value: int) -> int:
    return value + value
def traced1937_4(value: int) -> int:
    return value + 91


def decorate1937_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1937_4
def doubled1937_4(value: int) -> int:
    return value + value
def traced1937_5(value: int) -> int:
    return value + 98


def decorate1937_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1937_5
def doubled1937_5(value: int) -> int:
    return value + value
def expand1937_6(values):
    for item in values:
        if item > 8:
            yield item * 31


def collect1937_6(values):
    return list(expand1937_6(values))
def pair1937_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 15}
class Slot1937_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{22}"

    @property
    def doubled(self) -> int:
        return self.size * 2
