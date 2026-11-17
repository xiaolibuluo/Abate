
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize265_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair265_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 62}
class Slot265_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{69}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand265_3(values):
    for item in values:
        if item > 76:
            yield item * 67


def collect265_3(values):
    return list(expand265_3(values))
def expand265_4(values):
    for item in values:
        if item > 83:
            yield item * 17


def collect265_4(values):
    return list(expand265_4(values))
