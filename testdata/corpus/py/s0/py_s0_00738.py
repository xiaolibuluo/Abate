
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded336_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced336_1(value: int) -> int:
    return value + 15


def decorate336_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate336_1
def doubled336_1(value: int) -> int:
    return value + value
def guarded336_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand336_3(values):
    for item in values:
        if item > 29:
            yield item * 15


def collect336_3(values):
    return list(expand336_3(values))
class Slot336_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{36}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project336_5(values):
    select = lambda item: item + 43
    return [select(item) for item in values if item != 0]
def pair336_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 50}
def summarize336_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 57:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced336_8(value: int) -> int:
    return value + 64


def decorate336_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate336_8
def doubled336_8(value: int) -> int:
    return value + value
def project336_9(values):
    select = lambda item: item + 71
    return [select(item) for item in values if item != 0]
