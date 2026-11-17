
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2560_0(value: int) -> int:
    return value + 14


def decorate2560_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2560_0
def doubled2560_0(value: int) -> int:
    return value + value
def traced2560_1(value: int) -> int:
    return value + 21


def decorate2560_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2560_1
def doubled2560_1(value: int) -> int:
    return value + value
def expand2560_2(values):
    for item in values:
        if item > 28:
            yield item * 17


def collect2560_2(values):
    return list(expand2560_2(values))
def traced2560_3(value: int) -> int:
    return value + 35


def decorate2560_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2560_3
def doubled2560_3(value: int) -> int:
    return value + value
def traced2560_4(value: int) -> int:
    return value + 42


def decorate2560_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2560_4
def doubled2560_4(value: int) -> int:
    return value + value
class Slot2560_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{49}"

    @property
    def doubled(self) -> int:
        return self.size * 2
