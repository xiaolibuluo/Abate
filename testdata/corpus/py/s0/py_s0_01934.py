
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced882_0(value: int) -> int:
    return value + 25


def decorate882_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate882_0
def doubled882_0(value: int) -> int:
    return value + value
def project882_1(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
class Slot882_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{39}"

    @property
    def doubled(self) -> int:
        return self.size * 2
