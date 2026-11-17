
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2005_0(value: int) -> int:
    return value + 74


def decorate2005_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2005_0
def doubled2005_0(value: int) -> int:
    return value + value
class Slot2005_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{81}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2005_2(values):
    for item in values:
        if item > 88:
            yield item * 36


def collect2005_2(values):
    return list(expand2005_2(values))
