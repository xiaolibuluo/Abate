
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand30_0(values):
    for item in values:
        if item > 7:
            yield item * 36


def collect30_0(values):
    return list(expand30_0(values))
def expand30_1(values):
    for item in values:
        if item > 14:
            yield item * 47


def collect30_1(values):
    return list(expand30_1(values))
def summarize30_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded30_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot30_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{35}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced30_5(value: int) -> int:
    return value + 42


def decorate30_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate30_5
def doubled30_5(value: int) -> int:
    return value + value
def project30_6(values):
    select = lambda item: item + 49
    return [select(item) for item in values if item != 0]
