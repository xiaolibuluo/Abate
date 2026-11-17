
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot883_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{38}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced883_1(value: int) -> int:
    return value + 45


def decorate883_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate883_1
def doubled883_1(value: int) -> int:
    return value + value
