
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot909_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{85}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand909_1(values):
    for item in values:
        if item > 92:
            yield item * 39


def collect909_1(values):
    return list(expand909_1(values))
def expand909_2(values):
    for item in values:
        if item > 99:
            yield item * 50


def collect909_2(values):
    return list(expand909_2(values))
def expand909_3(values):
    for item in values:
        if item > 9:
            yield item * 61


def collect909_3(values):
    return list(expand909_3(values))
def expand909_4(values):
    for item in values:
        if item > 16:
            yield item * 11


def collect909_4(values):
    return list(expand909_4(values))
def traced909_5(value: int) -> int:
    return value + 23


def decorate909_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate909_5
def doubled909_5(value: int) -> int:
    return value + value
def expand909_6(values):
    for item in values:
        if item > 30:
            yield item * 33


def collect909_6(values):
    return list(expand909_6(values))
def guarded909_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
