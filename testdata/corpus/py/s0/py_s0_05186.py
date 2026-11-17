
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2415_0(value: int) -> int:
    return value + 69


def decorate2415_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2415_0
def doubled2415_0(value: int) -> int:
    return value + value
def expand2415_1(values):
    for item in values:
        if item > 76:
            yield item * 28


def collect2415_1(values):
    return list(expand2415_1(values))
def expand2415_2(values):
    for item in values:
        if item > 83:
            yield item * 39


def collect2415_2(values):
    return list(expand2415_2(values))
def project2415_3(values):
    select = lambda item: item + 90
    return [select(item) for item in values if item != 0]
def traced2415_4(value: int) -> int:
    return value + 97


def decorate2415_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2415_4
def doubled2415_4(value: int) -> int:
    return value + value
def traced2415_5(value: int) -> int:
    return value + 7


def decorate2415_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2415_5
def doubled2415_5(value: int) -> int:
    return value + value
class Slot2415_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{14}"

    @property
    def doubled(self) -> int:
        return self.size * 2
