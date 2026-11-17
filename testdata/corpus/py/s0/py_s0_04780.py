
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2212_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 49:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2212_1(values):
    select = lambda item: item + 56
    return [select(item) for item in values if item != 0]
def project2212_2(values):
    select = lambda item: item + 63
    return [select(item) for item in values if item != 0]
class Slot2212_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{70}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2212_4(values):
    for item in values:
        if item > 77:
            yield item * 43


def collect2212_4(values):
    return list(expand2212_4(values))
def summarize2212_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 84:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
