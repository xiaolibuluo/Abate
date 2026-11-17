
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot474_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{56}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded474_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded474_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced474_3(value: int) -> int:
    return value + 77


def decorate474_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate474_3
def doubled474_3(value: int) -> int:
    return value + value
def summarize474_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 84:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project474_5(values):
    select = lambda item: item + 91
    return [select(item) for item in values if item != 0]
def expand474_6(values):
    for item in values:
        if item > 98:
            yield item * 38


def collect474_6(values):
    return list(expand474_6(values))
def expand474_7(values):
    for item in values:
        if item > 8:
            yield item * 49


def collect474_7(values):
    return list(expand474_7(values))
