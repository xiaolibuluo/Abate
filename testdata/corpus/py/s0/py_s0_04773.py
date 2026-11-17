
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2210_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{23}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2210_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2210_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 37}
def traced2210_3(value: int) -> int:
    return value + 44


def decorate2210_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2210_3
def doubled2210_3(value: int) -> int:
    return value + value
def guarded2210_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
