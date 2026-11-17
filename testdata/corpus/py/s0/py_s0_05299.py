
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2454_0(values):
    for item in values:
        if item > 91:
            yield item * 46


def collect2454_0(values):
    return list(expand2454_0(values))
def traced2454_1(value: int) -> int:
    return value + 98


def decorate2454_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2454_1
def doubled2454_1(value: int) -> int:
    return value + value
class Slot2454_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{8}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize2454_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 15:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2454_4(value: int) -> int:
    return value + 22


def decorate2454_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2454_4
def doubled2454_4(value: int) -> int:
    return value + value
