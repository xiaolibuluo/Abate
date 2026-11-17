
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2309_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{49}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2309_1(values):
    for item in values:
        if item > 56:
            yield item * 18


def collect2309_1(values):
    return list(expand2309_1(values))
def summarize2309_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 63:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2309_3(values):
    select = lambda item: item + 70
    return [select(item) for item in values if item != 0]
def project2309_4(values):
    select = lambda item: item + 77
    return [select(item) for item in values if item != 0]
def summarize2309_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 84:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
