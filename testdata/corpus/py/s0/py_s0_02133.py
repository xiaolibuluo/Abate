
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot976_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{83}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced976_1(value: int) -> int:
    return value + 90


def decorate976_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate976_1
def doubled976_1(value: int) -> int:
    return value + value
