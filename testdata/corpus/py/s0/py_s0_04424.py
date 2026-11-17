
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2061_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{26}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2061_1(value: int) -> int:
    return value + 33


def decorate2061_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2061_1
def doubled2061_1(value: int) -> int:
    return value + value
class Slot2061_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot2061_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{47}"

    @property
    def doubled(self) -> int:
        return self.size * 2
