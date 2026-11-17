
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced218_0(value: int) -> int:
    return value + 26


def decorate218_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate218_0
def doubled218_0(value: int) -> int:
    return value + value
def guarded218_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot218_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot218_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{47}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot218_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{54}"

    @property
    def doubled(self) -> int:
        return self.size * 2
