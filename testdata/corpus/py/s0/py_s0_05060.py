
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2353_0(values):
    select = lambda item: item + 39
    return [select(item) for item in values if item != 0]
def summarize2353_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 46:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2353_2(values):
    for item in values:
        if item > 53:
            yield item * 32


def collect2353_2(values):
    return list(expand2353_2(values))
def expand2353_3(values):
    for item in values:
        if item > 60:
            yield item * 43


def collect2353_3(values):
    return list(expand2353_3(values))
class Slot2353_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{67}"

    @property
    def doubled(self) -> int:
        return self.size * 2
