
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot960_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{69}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot960_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{76}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize960_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 83:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project960_3(values):
    select = lambda item: item + 90
    return [select(item) for item in values if item != 0]
def project960_4(values):
    select = lambda item: item + 97
    return [select(item) for item in values if item != 0]
def summarize960_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 7:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand960_6(values):
    for item in values:
        if item > 14:
            yield item * 24


def collect960_6(values):
    return list(expand960_6(values))
class Slot960_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{21}"

    @property
    def doubled(self) -> int:
        return self.size * 2
