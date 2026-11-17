
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2108_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2108_1(values):
    for item in values:
        if item > 62:
            yield item * 14


def collect2108_1(values):
    return list(expand2108_1(values))
class Slot2108_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{69}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize2108_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 76:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
