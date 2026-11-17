
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1190_0(values):
    select = lambda item: item + 52
    return [select(item) for item in values if item != 0]
def summarize1190_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 59:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1190_2(values):
    for item in values:
        if item > 66:
            yield item * 65


def collect1190_2(values):
    return list(expand1190_2(values))
class Slot1190_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{73}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1190_4(value: int) -> int:
    return value + 80


def decorate1190_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1190_4
def doubled1190_4(value: int) -> int:
    return value + value
