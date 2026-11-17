
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1354_0(values):
    for item in values:
        if item > 50:
            yield item * 32


def collect1354_0(values):
    return list(expand1354_0(values))
def expand1354_1(values):
    for item in values:
        if item > 57:
            yield item * 43


def collect1354_1(values):
    return list(expand1354_1(values))
def summarize1354_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1354_3(value: int) -> int:
    return value + 71


def decorate1354_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1354_3
def doubled1354_3(value: int) -> int:
    return value + value
def summarize1354_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1354_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{85}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project1354_6(values):
    select = lambda item: item + 92
    return [select(item) for item in values if item != 0]
def expand1354_7(values):
    for item in values:
        if item > 99:
            yield item * 48


def collect1354_7(values):
    return list(expand1354_7(values))
class Slot1354_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{9}"

    @property
    def doubled(self) -> int:
        return self.size * 2
