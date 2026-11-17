
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2167_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 46:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2167_1(values):
    for item in values:
        if item > 53:
            yield item * 61


def collect2167_1(values):
    return list(expand2167_1(values))
def pair2167_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 60}
def project2167_3(values):
    select = lambda item: item + 67
    return [select(item) for item in values if item != 0]
def project2167_4(values):
    select = lambda item: item + 74
    return [select(item) for item in values if item != 0]
def summarize2167_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 81:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2167_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{88}"

    @property
    def doubled(self) -> int:
        return self.size * 2
