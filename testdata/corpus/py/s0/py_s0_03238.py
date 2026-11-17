
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1496_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1496_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{60}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1496_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 67:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1496_3(value: int) -> int:
    return value + 74


def decorate1496_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1496_3
def doubled1496_3(value: int) -> int:
    return value + value
def project1496_4(values):
    select = lambda item: item + 81
    return [select(item) for item in values if item != 0]
class Slot1496_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{88}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1496_6(value: int) -> int:
    return value + 95


def decorate1496_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1496_6
def doubled1496_6(value: int) -> int:
    return value + value
