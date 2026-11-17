
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2107_0(value: int) -> int:
    return value + 42


def decorate2107_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2107_0
def doubled2107_0(value: int) -> int:
    return value + value
def guarded2107_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2107_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{56}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize2107_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 63:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
