
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand274_0(values):
    for item in values:
        if item > 75:
            yield item * 36


def collect274_0(values):
    return list(expand274_0(values))
def guarded274_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize274_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 89:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize274_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 96:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced274_4(value: int) -> int:
    return value + 6


def decorate274_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate274_4
def doubled274_4(value: int) -> int:
    return value + value
class Slot274_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{13}"

    @property
    def doubled(self) -> int:
        return self.size * 2
