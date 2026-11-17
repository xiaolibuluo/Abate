
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded288_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded288_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced288_2(value: int) -> int:
    return value + 77


def decorate288_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate288_2
def doubled288_2(value: int) -> int:
    return value + value
def pair288_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 84}
def traced288_4(value: int) -> int:
    return value + 91


def decorate288_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate288_4
def doubled288_4(value: int) -> int:
    return value + value
def summarize288_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 98:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot288_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{8}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded288_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced288_8(value: int) -> int:
    return value + 22


def decorate288_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate288_8
def doubled288_8(value: int) -> int:
    return value + value
