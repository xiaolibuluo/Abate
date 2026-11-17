
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1407_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 60:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1407_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 67:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1407_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1407_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 81:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1407_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 88}
def traced1407_5(value: int) -> int:
    return value + 95


def decorate1407_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1407_5
def doubled1407_5(value: int) -> int:
    return value + value
def traced1407_6(value: int) -> int:
    return value + 5


def decorate1407_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1407_6
def doubled1407_6(value: int) -> int:
    return value + value
class Slot1407_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{12}"

    @property
    def doubled(self) -> int:
        return self.size * 2
