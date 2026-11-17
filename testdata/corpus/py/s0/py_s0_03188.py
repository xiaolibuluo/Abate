
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1475_0(value: int) -> int:
    return value + 71


def decorate1475_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1475_0
def doubled1475_0(value: int) -> int:
    return value + value
def project1475_1(values):
    select = lambda item: item + 78
    return [select(item) for item in values if item != 0]
class Slot1475_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{85}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1475_3(values):
    for item in values:
        if item > 92:
            yield item * 58


def collect1475_3(values):
    return list(expand1475_3(values))
