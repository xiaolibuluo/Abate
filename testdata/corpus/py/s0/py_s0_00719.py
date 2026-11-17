
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize328_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 98:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded328_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced328_2(value: int) -> int:
    return value + 15


def decorate328_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate328_2
def doubled328_2(value: int) -> int:
    return value + value
def traced328_3(value: int) -> int:
    return value + 22


def decorate328_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate328_3
def doubled328_3(value: int) -> int:
    return value + value
def expand328_4(values):
    for item in values:
        if item > 29:
            yield item * 31


def collect328_4(values):
    return list(expand328_4(values))
def guarded328_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand328_6(values):
    for item in values:
        if item > 43:
            yield item * 53


def collect328_6(values):
    return list(expand328_6(values))
class Slot328_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{50}"

    @property
    def doubled(self) -> int:
        return self.size * 2
