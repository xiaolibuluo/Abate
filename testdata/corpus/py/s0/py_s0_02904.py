
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1332_0(value: int) -> int:
    return value + 55


def decorate1332_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1332_0
def doubled1332_0(value: int) -> int:
    return value + value
def traced1332_1(value: int) -> int:
    return value + 62


def decorate1332_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1332_1
def doubled1332_1(value: int) -> int:
    return value + value
class Slot1332_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{69}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1332_3(values):
    for item in values:
        if item > 76:
            yield item * 33


def collect1332_3(values):
    return list(expand1332_3(values))
