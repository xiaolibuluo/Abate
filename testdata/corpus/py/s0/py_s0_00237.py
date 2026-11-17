
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project107_0(values):
    select = lambda item: item + 38
    return [select(item) for item in values if item != 0]
class Slot107_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{45}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced107_2(value: int) -> int:
    return value + 52


def decorate107_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate107_2
def doubled107_2(value: int) -> int:
    return value + value
