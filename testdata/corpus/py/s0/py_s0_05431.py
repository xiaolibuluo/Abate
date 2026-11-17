
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2506_0(values):
    for item in values:
        if item > 88:
            yield item * 44


def collect2506_0(values):
    return list(expand2506_0(values))
def traced2506_1(value: int) -> int:
    return value + 95


def decorate2506_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2506_1
def doubled2506_1(value: int) -> int:
    return value + value
def guarded2506_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2506_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 12:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2506_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2506_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{26}"

    @property
    def doubled(self) -> int:
        return self.size * 2
