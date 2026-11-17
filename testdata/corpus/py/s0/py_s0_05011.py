
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2326_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{76}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2326_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2326_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 90:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2326_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2326_4(value: int) -> int:
    return value + 7


def decorate2326_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2326_4
def doubled2326_4(value: int) -> int:
    return value + value
def traced2326_5(value: int) -> int:
    return value + 14


def decorate2326_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2326_5
def doubled2326_5(value: int) -> int:
    return value + value
def project2326_6(values):
    select = lambda item: item + 21
    return [select(item) for item in values if item != 0]
class Slot2326_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{28}"

    @property
    def doubled(self) -> int:
        return self.size * 2
