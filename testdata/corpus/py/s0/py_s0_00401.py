
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded187_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair187_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 18}
def guarded187_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project187_3(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
def traced187_4(value: int) -> int:
    return value + 39


def decorate187_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate187_4
def doubled187_4(value: int) -> int:
    return value + value
def summarize187_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 46:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot187_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{53}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize187_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 60:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project187_8(values):
    select = lambda item: item + 67
    return [select(item) for item in values if item != 0]
