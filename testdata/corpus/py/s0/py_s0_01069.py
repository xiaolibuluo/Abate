
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand495_0(values):
    for item in values:
        if item > 38:
            yield item * 58


def collect495_0(values):
    return list(expand495_0(values))
class Slot495_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{45}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize495_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 52:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
