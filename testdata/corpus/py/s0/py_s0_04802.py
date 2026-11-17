
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2223_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{95}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2223_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2223_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 12}
def traced2223_3(value: int) -> int:
    return value + 19


def decorate2223_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2223_3
def doubled2223_3(value: int) -> int:
    return value + value
def traced2223_4(value: int) -> int:
    return value + 26


def decorate2223_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2223_4
def doubled2223_4(value: int) -> int:
    return value + value
def guarded2223_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2223_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 40:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
