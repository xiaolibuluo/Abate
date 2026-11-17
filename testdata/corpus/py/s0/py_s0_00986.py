
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand456_0(values):
    for item in values:
        if item > 16:
            yield item * 29


def collect456_0(values):
    return list(expand456_0(values))
class Slot456_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{23}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize456_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair456_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 37}
