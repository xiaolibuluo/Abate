
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2104_0(values):
    for item in values:
        if item > 100:
            yield item * 36


def collect2104_0(values):
    return list(expand2104_0(values))
def traced2104_1(value: int) -> int:
    return value + 10


def decorate2104_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2104_1
def doubled2104_1(value: int) -> int:
    return value + value
class Slot2104_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{17}"

    @property
    def doubled(self) -> int:
        return self.size * 2
