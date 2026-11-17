
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot421_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{46}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project421_1(values):
    select = lambda item: item + 53
    return [select(item) for item in values if item != 0]
class Slot421_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{60}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced421_3(value: int) -> int:
    return value + 67


def decorate421_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate421_3
def doubled421_3(value: int) -> int:
    return value + value
def expand421_4(values):
    for item in values:
        if item > 74:
            yield item * 11


def collect421_4(values):
    return list(expand421_4(values))
def guarded421_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize421_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 88:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
