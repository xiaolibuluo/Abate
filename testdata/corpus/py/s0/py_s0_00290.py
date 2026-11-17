
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced131_0(value: int) -> int:
    return value + 59


def decorate131_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate131_0
def doubled131_0(value: int) -> int:
    return value + value
def summarize131_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 66:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot131_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{73}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand131_3(values):
    for item in values:
        if item > 80:
            yield item * 44


def collect131_3(values):
    return list(expand131_3(values))
def summarize131_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 87:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded131_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded131_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
