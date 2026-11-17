
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair252_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 80}
def traced252_1(value: int) -> int:
    return value + 87


def decorate252_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate252_1
def doubled252_1(value: int) -> int:
    return value + value
class Slot252_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{94}"

    @property
    def doubled(self) -> int:
        return self.size * 2
