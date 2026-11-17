
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1970_0(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
def expand1970_1(values):
    for item in values:
        if item > 14:
            yield item * 24


def collect1970_1(values):
    return list(expand1970_1(values))
def guarded1970_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1970_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{28}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1970_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 35:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1970_5(value: int) -> int:
    return value + 42


def decorate1970_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1970_5
def doubled1970_5(value: int) -> int:
    return value + value
