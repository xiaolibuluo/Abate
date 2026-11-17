
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize168_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand168_1(values):
    for item in values:
        if item > 62:
            yield item * 37


def collect168_1(values):
    return list(expand168_1(values))
def summarize168_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 69:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project168_3(values):
    select = lambda item: item + 76
    return [select(item) for item in values if item != 0]
class Slot168_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{83}"

    @property
    def doubled(self) -> int:
        return self.size * 2
