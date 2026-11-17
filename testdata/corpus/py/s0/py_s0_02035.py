
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced928_0(value: int) -> int:
    return value + 41


def decorate928_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate928_0
def doubled928_0(value: int) -> int:
    return value + value
class Slot928_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{48}"

    @property
    def doubled(self) -> int:
        return self.size * 2
