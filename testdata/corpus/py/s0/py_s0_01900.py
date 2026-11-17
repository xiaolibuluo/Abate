
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded866_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded866_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced866_2(value: int) -> int:
    return value + 25


def decorate866_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate866_2
def doubled866_2(value: int) -> int:
    return value + value
def pair866_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 32}
def project866_4(values):
    select = lambda item: item + 39
    return [select(item) for item in values if item != 0]
class Slot866_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{46}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize866_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 53:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot866_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{60}"

    @property
    def doubled(self) -> int:
        return self.size * 2
