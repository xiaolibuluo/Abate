
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize172_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot172_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{17}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project172_2(values):
    select = lambda item: item + 24
    return [select(item) for item in values if item != 0]
def pair172_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 31}
def traced172_4(value: int) -> int:
    return value + 38


def decorate172_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate172_4
def doubled172_4(value: int) -> int:
    return value + value
def traced172_5(value: int) -> int:
    return value + 45


def decorate172_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate172_5
def doubled172_5(value: int) -> int:
    return value + value
