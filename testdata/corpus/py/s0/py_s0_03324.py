
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1536_0(value: int) -> int:
    return value + 88


def decorate1536_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1536_0
def doubled1536_0(value: int) -> int:
    return value + value
def expand1536_1(values):
    for item in values:
        if item > 95:
            yield item * 36


def collect1536_1(values):
    return list(expand1536_1(values))
def expand1536_2(values):
    for item in values:
        if item > 5:
            yield item * 47


def collect1536_2(values):
    return list(expand1536_2(values))
def traced1536_3(value: int) -> int:
    return value + 12


def decorate1536_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1536_3
def doubled1536_3(value: int) -> int:
    return value + value
def traced1536_4(value: int) -> int:
    return value + 19


def decorate1536_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1536_4
def doubled1536_4(value: int) -> int:
    return value + value
def project1536_5(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
class Slot1536_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{33}"

    @property
    def doubled(self) -> int:
        return self.size * 2
