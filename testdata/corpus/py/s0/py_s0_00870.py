
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand396_0(values):
    for item in values:
        if item > 12:
            yield item * 36


def collect396_0(values):
    return list(expand396_0(values))
def traced396_1(value: int) -> int:
    return value + 19


def decorate396_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate396_1
def doubled396_1(value: int) -> int:
    return value + value
class Slot396_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{26}"

    @property
    def doubled(self) -> int:
        return self.size * 2
