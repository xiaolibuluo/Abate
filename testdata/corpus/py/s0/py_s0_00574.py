
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize262_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 16:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced262_1(value: int) -> int:
    return value + 23


def decorate262_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate262_1
def doubled262_1(value: int) -> int:
    return value + value
class Slot262_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{30}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project262_3(values):
    select = lambda item: item + 37
    return [select(item) for item in values if item != 0]
