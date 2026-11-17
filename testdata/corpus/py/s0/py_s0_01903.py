
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded868_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded868_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project868_2(values):
    select = lambda item: item + 51
    return [select(item) for item in values if item != 0]
def project868_3(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
def expand868_4(values):
    for item in values:
        if item > 65:
            yield item * 29


def collect868_4(values):
    return list(expand868_4(values))
class Slot868_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{72}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize868_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 79:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair868_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 86}
def expand868_8(values):
    for item in values:
        if item > 93:
            yield item * 12


def collect868_8(values):
    return list(expand868_8(values))
def traced868_9(value: int) -> int:
    return value + 100


def decorate868_9(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate868_9
def doubled868_9(value: int) -> int:
    return value + value
