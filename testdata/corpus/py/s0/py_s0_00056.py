
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded28_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair28_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 85}
def project28_2(values):
    select = lambda item: item + 92
    return [select(item) for item in values if item != 0]
class Slot28_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{99}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand28_4(values):
    for item in values:
        if item > 9:
            yield item * 66


def collect28_4(values):
    return list(expand28_4(values))
def traced28_5(value: int) -> int:
    return value + 16


def decorate28_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate28_5
def doubled28_5(value: int) -> int:
    return value + value
def guarded28_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
