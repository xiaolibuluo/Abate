
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1052_0(values):
    for item in values:
        if item > 101:
            yield item * 53


def collect1052_0(values):
    return list(expand1052_0(values))
def expand1052_1(values):
    for item in values:
        if item > 11:
            yield item * 64


def collect1052_1(values):
    return list(expand1052_1(values))
def guarded1052_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1052_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 25:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1052_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{32}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1052_5(value: int) -> int:
    return value + 39


def decorate1052_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1052_5
def doubled1052_5(value: int) -> int:
    return value + value
def traced1052_6(value: int) -> int:
    return value + 46


def decorate1052_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1052_6
def doubled1052_6(value: int) -> int:
    return value + value
