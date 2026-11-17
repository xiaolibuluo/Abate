
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2561_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 27:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2561_1(values):
    for item in values:
        if item > 34:
            yield item * 13


def collect2561_1(values):
    return list(expand2561_1(values))
def guarded2561_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2561_3(value: int) -> int:
    return value + 48


def decorate2561_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2561_3
def doubled2561_3(value: int) -> int:
    return value + value
class Slot2561_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{55}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot2561_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{62}"

    @property
    def doubled(self) -> int:
        return self.size * 2
