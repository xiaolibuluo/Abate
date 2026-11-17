
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand261_0(values):
    for item in values:
        if item > 100:
            yield item * 67


def collect261_0(values):
    return list(expand261_0(values))
class Slot261_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{10}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize261_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 17:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project261_3(values):
    select = lambda item: item + 24
    return [select(item) for item in values if item != 0]
def traced261_4(value: int) -> int:
    return value + 31


def decorate261_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate261_4
def doubled261_4(value: int) -> int:
    return value + value
