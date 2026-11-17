
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2014_0(values):
    select = lambda item: item + 94
    return [select(item) for item in values if item != 0]
def summarize2014_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 101:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2014_2(values):
    for item in values:
        if item > 11:
            yield item * 38


def collect2014_2(values):
    return list(expand2014_2(values))
def expand2014_3(values):
    for item in values:
        if item > 18:
            yield item * 49


def collect2014_3(values):
    return list(expand2014_3(values))
class Slot2014_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{25}"

    @property
    def doubled(self) -> int:
        return self.size * 2
