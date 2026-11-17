
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1662_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{77}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1662_1(values):
    for item in values:
        if item > 84:
            yield item * 64


def collect1662_1(values):
    return list(expand1662_1(values))
def summarize1662_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 91:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
