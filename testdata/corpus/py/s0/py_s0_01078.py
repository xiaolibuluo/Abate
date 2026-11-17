
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced499_0(value: int) -> int:
    return value + 90


def decorate499_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate499_0
def doubled499_0(value: int) -> int:
    return value + value
class Slot499_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced499_2(value: int) -> int:
    return value + 7


def decorate499_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate499_2
def doubled499_2(value: int) -> int:
    return value + value
