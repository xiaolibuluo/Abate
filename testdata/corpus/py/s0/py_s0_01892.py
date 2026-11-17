
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded861_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair861_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 50}
class Slot861_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{57}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize861_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot861_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{71}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced861_5(value: int) -> int:
    return value + 78


def decorate861_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate861_5
def doubled861_5(value: int) -> int:
    return value + value
def summarize861_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 85:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
