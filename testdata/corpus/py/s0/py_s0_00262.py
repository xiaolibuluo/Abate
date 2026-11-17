
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand117_0(values):
    for item in values:
        if item > 71:
            yield item * 35


def collect117_0(values):
    return list(expand117_0(values))
def summarize117_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced117_2(value: int) -> int:
    return value + 85


def decorate117_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate117_2
def doubled117_2(value: int) -> int:
    return value + value
class Slot117_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{92}"

    @property
    def doubled(self) -> int:
        return self.size * 2
