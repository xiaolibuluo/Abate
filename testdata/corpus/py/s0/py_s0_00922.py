
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced423_0(value: int) -> int:
    return value + 72


def decorate423_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate423_0
def doubled423_0(value: int) -> int:
    return value + value
def guarded423_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize423_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 86:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot423_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{93}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded423_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded423_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand423_6(values):
    for item in values:
        if item > 17:
            yield item * 47


def collect423_6(values):
    return list(expand423_6(values))
