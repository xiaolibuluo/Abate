
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand213_0(values):
    for item in values:
        if item > 58:
            yield item * 36


def collect213_0(values):
    return list(expand213_0(values))
class Slot213_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot213_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{72}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair213_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 79}
def summarize213_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 86:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
