
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1237_0(values):
    for item in values:
        if item > 81:
            yield item * 67


def collect1237_0(values):
    return list(expand1237_0(values))
def expand1237_1(values):
    for item in values:
        if item > 88:
            yield item * 17


def collect1237_1(values):
    return list(expand1237_1(values))
def pair1237_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 95}
def summarize1237_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 5:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1237_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{12}"

    @property
    def doubled(self) -> int:
        return self.size * 2
