
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2124_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 69:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2124_1(values):
    for item in values:
        if item > 76:
            yield item * 65


def collect2124_1(values):
    return list(expand2124_1(values))
class Slot2124_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{83}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize2124_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 90:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2124_4(value: int) -> int:
    return value + 97


def decorate2124_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2124_4
def doubled2124_4(value: int) -> int:
    return value + value
class Slot2124_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{7}"

    @property
    def doubled(self) -> int:
        return self.size * 2
