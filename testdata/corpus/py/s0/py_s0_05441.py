
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2514_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 95:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2514_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{5}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2514_2(values):
    for item in values:
        if item > 12:
            yield item * 61


def collect2514_2(values):
    return list(expand2514_2(values))
def summarize2514_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 19:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2514_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{26}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair2514_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 33}
def traced2514_6(value: int) -> int:
    return value + 40


def decorate2514_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2514_6
def doubled2514_6(value: int) -> int:
    return value + value
