
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1201_0(values):
    select = lambda item: item + 98
    return [select(item) for item in values if item != 0]
def traced1201_1(value: int) -> int:
    return value + 8


def decorate1201_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1201_1
def doubled1201_1(value: int) -> int:
    return value + value
def traced1201_2(value: int) -> int:
    return value + 15


def decorate1201_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1201_2
def doubled1201_2(value: int) -> int:
    return value + value
class Slot1201_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{22}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1201_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 29:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1201_5(values):
    for item in values:
        if item > 36:
            yield item * 53


def collect1201_5(values):
    return list(expand1201_5(values))
def traced1201_6(value: int) -> int:
    return value + 43


def decorate1201_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1201_6
def doubled1201_6(value: int) -> int:
    return value + value
