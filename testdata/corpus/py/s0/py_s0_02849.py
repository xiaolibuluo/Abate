
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1303_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 66}
def guarded1303_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1303_2(value: int) -> int:
    return value + 80


def decorate1303_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1303_2
def doubled1303_2(value: int) -> int:
    return value + value
class Slot1303_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{87}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1303_4(value: int) -> int:
    return value + 94


def decorate1303_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1303_4
def doubled1303_4(value: int) -> int:
    return value + value
def summarize1303_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 101:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1303_6(values):
    for item in values:
        if item > 11:
            yield item * 46


def collect1303_6(values):
    return list(expand1303_6(values))
def guarded1303_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
