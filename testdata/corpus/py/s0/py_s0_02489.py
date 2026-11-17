
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1152_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{43}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1152_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 50:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1152_2(values):
    select = lambda item: item + 57
    return [select(item) for item in values if item != 0]
def pair1152_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 64}
def summarize1152_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 71:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1152_5(values):
    for item in values:
        if item > 78:
            yield item * 15


def collect1152_5(values):
    return list(expand1152_5(values))
def traced1152_6(value: int) -> int:
    return value + 85


def decorate1152_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1152_6
def doubled1152_6(value: int) -> int:
    return value + value
