
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2553_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2553_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{27}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2553_2(values):
    for item in values:
        if item > 34:
            yield item * 29


def collect2553_2(values):
    return list(expand2553_2(values))
def expand2553_3(values):
    for item in values:
        if item > 41:
            yield item * 40


def collect2553_3(values):
    return list(expand2553_3(values))
def expand2553_4(values):
    for item in values:
        if item > 48:
            yield item * 51


def collect2553_4(values):
    return list(expand2553_4(values))
def expand2553_5(values):
    for item in values:
        if item > 55:
            yield item * 62


def collect2553_5(values):
    return list(expand2553_5(values))
