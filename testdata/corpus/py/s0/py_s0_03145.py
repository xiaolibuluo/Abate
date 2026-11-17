
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1450_0(values):
    for item in values:
        if item > 37:
            yield item * 33


def collect1450_0(values):
    return list(expand1450_0(values))
def traced1450_1(value: int) -> int:
    return value + 44


def decorate1450_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1450_1
def doubled1450_1(value: int) -> int:
    return value + value
def summarize1450_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 51:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1450_3(value: int) -> int:
    return value + 58


def decorate1450_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1450_3
def doubled1450_3(value: int) -> int:
    return value + value
class Slot1450_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1450_5(value: int) -> int:
    return value + 72


def decorate1450_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1450_5
def doubled1450_5(value: int) -> int:
    return value + value
def summarize1450_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 79:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
