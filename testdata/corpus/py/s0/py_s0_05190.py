
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2418_0(values):
    for item in values:
        if item > 11:
            yield item * 38


def collect2418_0(values):
    return list(expand2418_0(values))
def expand2418_1(values):
    for item in values:
        if item > 18:
            yield item * 49


def collect2418_1(values):
    return list(expand2418_1(values))
def summarize2418_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 25:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2418_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 32}
class Slot2418_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{39}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair2418_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 46}
def summarize2418_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 53:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
