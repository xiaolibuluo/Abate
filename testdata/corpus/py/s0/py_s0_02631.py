
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1203_0(value: int) -> int:
    return value + 27


def decorate1203_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1203_0
def doubled1203_0(value: int) -> int:
    return value + value
def summarize1203_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 34:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1203_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{41}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1203_3(value: int) -> int:
    return value + 48


def decorate1203_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1203_3
def doubled1203_3(value: int) -> int:
    return value + value
def summarize1203_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1203_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1203_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 69:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
