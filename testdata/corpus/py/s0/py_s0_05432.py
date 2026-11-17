
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2507_0(values):
    for item in values:
        if item > 101:
            yield item * 51


def collect2507_0(values):
    return list(expand2507_0(values))
class Slot2507_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{11}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot2507_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{18}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize2507_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 25:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2507_4(value: int) -> int:
    return value + 32


def decorate2507_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2507_4
def doubled2507_4(value: int) -> int:
    return value + value
