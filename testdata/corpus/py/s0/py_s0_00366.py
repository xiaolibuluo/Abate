
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize166_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 29:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot166_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{36}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced166_2(value: int) -> int:
    return value + 43


def decorate166_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate166_2
def doubled166_2(value: int) -> int:
    return value + value
def project166_3(values):
    select = lambda item: item + 50
    return [select(item) for item in values if item != 0]
def expand166_4(values):
    for item in values:
        if item > 57:
            yield item * 56


def collect166_4(values):
    return list(expand166_4(values))
