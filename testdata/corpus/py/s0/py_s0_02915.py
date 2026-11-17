
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1336_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{10}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1336_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 17:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1336_2(values):
    for item in values:
        if item > 24:
            yield item * 50


def collect1336_2(values):
    return list(expand1336_2(values))
def expand1336_3(values):
    for item in values:
        if item > 31:
            yield item * 61


def collect1336_3(values):
    return list(expand1336_3(values))
class Slot1336_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{38}"

    @property
    def doubled(self) -> int:
        return self.size * 2
