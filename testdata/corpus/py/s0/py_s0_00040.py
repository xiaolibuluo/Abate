
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot24_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{26}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded24_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot24_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand24_3(values):
    for item in values:
        if item > 47:
            yield item * 27


def collect24_3(values):
    return list(expand24_3(values))
def traced24_4(value: int) -> int:
    return value + 54


def decorate24_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate24_4
def doubled24_4(value: int) -> int:
    return value + value
def summarize24_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 61:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
