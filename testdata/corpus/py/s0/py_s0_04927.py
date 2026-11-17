
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2284_0(values):
    for item in values:
        if item > 15:
            yield item * 15


def collect2284_0(values):
    return list(expand2284_0(values))
def guarded2284_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2284_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 29:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2284_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 36:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2284_4(values):
    for item in values:
        if item > 43:
            yield item * 59


def collect2284_4(values):
    return list(expand2284_4(values))
def summarize2284_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 50:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2284_6(value: int) -> int:
    return value + 57


def decorate2284_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2284_6
def doubled2284_6(value: int) -> int:
    return value + value
def summarize2284_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2284_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{71}"

    @property
    def doubled(self) -> int:
        return self.size * 2
