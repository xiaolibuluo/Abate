
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded902_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced902_1(value: int) -> int:
    return value + 98


def decorate902_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate902_1
def doubled902_1(value: int) -> int:
    return value + value
def pair902_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 8}
class Slot902_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{15}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded902_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced902_5(value: int) -> int:
    return value + 29


def decorate902_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate902_5
def doubled902_5(value: int) -> int:
    return value + value
