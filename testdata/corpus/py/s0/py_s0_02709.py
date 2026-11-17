
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1239_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1239_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{17}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1239_2(value: int) -> int:
    return value + 24


def decorate1239_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1239_2
def doubled1239_2(value: int) -> int:
    return value + value
class Slot1239_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{31}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1239_4(value: int) -> int:
    return value + 38


def decorate1239_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1239_4
def doubled1239_4(value: int) -> int:
    return value + value
