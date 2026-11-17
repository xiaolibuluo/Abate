
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project864_0(values):
    select = lambda item: item + 82
    return [select(item) for item in values if item != 0]
def summarize864_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 89:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize864_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 96:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced864_3(value: int) -> int:
    return value + 6


def decorate864_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate864_3
def doubled864_3(value: int) -> int:
    return value + value
def summarize864_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot864_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{20}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded864_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
