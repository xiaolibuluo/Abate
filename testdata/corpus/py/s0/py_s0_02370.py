
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1092_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1092_1(values):
    for item in values:
        if item > 46:
            yield item * 39


def collect1092_1(values):
    return list(expand1092_1(values))
def traced1092_2(value: int) -> int:
    return value + 53


def decorate1092_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1092_2
def doubled1092_2(value: int) -> int:
    return value + value
class Slot1092_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{60}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1092_4(value: int) -> int:
    return value + 67


def decorate1092_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1092_4
def doubled1092_4(value: int) -> int:
    return value + value
