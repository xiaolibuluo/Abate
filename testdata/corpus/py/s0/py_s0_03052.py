
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1403_0(values):
    select = lambda item: item + 8
    return [select(item) for item in values if item != 0]
def pair1403_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 15}
def guarded1403_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1403_3(values):
    for item in values:
        if item > 29:
            yield item * 42


def collect1403_3(values):
    return list(expand1403_3(values))
class Slot1403_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{36}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1403_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1403_6(value: int) -> int:
    return value + 50


def decorate1403_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1403_6
def doubled1403_6(value: int) -> int:
    return value + value
def summarize1403_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 57:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1403_8(values):
    select = lambda item: item + 64
    return [select(item) for item in values if item != 0]
