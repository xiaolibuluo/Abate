
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1871_0(value: int) -> int:
    return value + 78


def decorate1871_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1871_0
def doubled1871_0(value: int) -> int:
    return value + value
def project1871_1(values):
    select = lambda item: item + 85
    return [select(item) for item in values if item != 0]
class Slot1871_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{92}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1871_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 99:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1871_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 9:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
