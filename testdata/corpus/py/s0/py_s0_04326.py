
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2016_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2016_1(value: int) -> int:
    return value + 30


def decorate2016_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2016_1
def doubled2016_1(value: int) -> int:
    return value + value
def project2016_2(values):
    select = lambda item: item + 37
    return [select(item) for item in values if item != 0]
def summarize2016_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 44:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2016_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{51}"

    @property
    def doubled(self) -> int:
        return self.size * 2
