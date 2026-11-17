
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1346_0(value: int) -> int:
    return value + 43


def decorate1346_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1346_0
def doubled1346_0(value: int) -> int:
    return value + value
def traced1346_1(value: int) -> int:
    return value + 50


def decorate1346_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1346_1
def doubled1346_1(value: int) -> int:
    return value + value
def guarded1346_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1346_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1346_4(value: int) -> int:
    return value + 71


def decorate1346_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1346_4
def doubled1346_4(value: int) -> int:
    return value + value
class Slot1346_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{78}"

    @property
    def doubled(self) -> int:
        return self.size * 2
