
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot965_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{37}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand965_1(values):
    for item in values:
        if item > 44:
            yield item * 65


def collect965_1(values):
    return list(expand965_1(values))
class Slot965_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{51}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project965_3(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
def expand965_4(values):
    for item in values:
        if item > 65:
            yield item * 37


def collect965_4(values):
    return list(expand965_4(values))
def traced965_5(value: int) -> int:
    return value + 72


def decorate965_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate965_5
def doubled965_5(value: int) -> int:
    return value + value
def guarded965_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot965_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{86}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project965_8(values):
    select = lambda item: item + 93
    return [select(item) for item in values if item != 0]
