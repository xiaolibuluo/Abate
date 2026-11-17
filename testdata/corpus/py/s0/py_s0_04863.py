
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2250_0(value: int) -> int:
    return value + 58


def decorate2250_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2250_0
def doubled2250_0(value: int) -> int:
    return value + value
class Slot2250_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2250_2(value: int) -> int:
    return value + 72


def decorate2250_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2250_2
def doubled2250_2(value: int) -> int:
    return value + value
