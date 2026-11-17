
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot475_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{69}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand475_1(values):
    for item in values:
        if item > 76:
            yield item * 51


def collect475_1(values):
    return list(expand475_1(values))
class Slot475_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{83}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project475_3(values):
    select = lambda item: item + 90
    return [select(item) for item in values if item != 0]
def expand475_4(values):
    for item in values:
        if item > 97:
            yield item * 23


def collect475_4(values):
    return list(expand475_4(values))
def expand475_5(values):
    for item in values:
        if item > 7:
            yield item * 34


def collect475_5(values):
    return list(expand475_5(values))
def summarize475_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize475_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
