
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2449_0(value: int) -> int:
    return value + 26


def decorate2449_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2449_0
def doubled2449_0(value: int) -> int:
    return value + value
def summarize2449_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 33:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2449_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2449_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
