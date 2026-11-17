
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize613_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot613_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{27}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand613_2(values):
    for item in values:
        if item > 34:
            yield item * 52


def collect613_2(values):
    return list(expand613_2(values))
def expand613_3(values):
    for item in values:
        if item > 41:
            yield item * 63


def collect613_3(values):
    return list(expand613_3(values))
def traced613_4(value: int) -> int:
    return value + 48


def decorate613_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate613_4
def doubled613_4(value: int) -> int:
    return value + value
def summarize613_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
