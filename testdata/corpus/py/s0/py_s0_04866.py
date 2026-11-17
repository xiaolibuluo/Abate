
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2252_0(values):
    for item in values:
        if item > 84:
            yield item * 35


def collect2252_0(values):
    return list(expand2252_0(values))
def expand2252_1(values):
    for item in values:
        if item > 91:
            yield item * 46


def collect2252_1(values):
    return list(expand2252_1(values))
class Slot2252_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{98}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2252_3(value: int) -> int:
    return value + 8


def decorate2252_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2252_3
def doubled2252_3(value: int) -> int:
    return value + value
