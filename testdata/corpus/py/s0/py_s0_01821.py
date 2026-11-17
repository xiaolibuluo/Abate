
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot834_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{80}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced834_1(value: int) -> int:
    return value + 87


def decorate834_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate834_1
def doubled834_1(value: int) -> int:
    return value + value
