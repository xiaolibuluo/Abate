
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1567_0(value: int) -> int:
    return value + 6


def decorate1567_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1567_0
def doubled1567_0(value: int) -> int:
    return value + value
class Slot1567_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{13}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1567_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{20}"

    @property
    def doubled(self) -> int:
        return self.size * 2
