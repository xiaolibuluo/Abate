
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize254_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 9:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot254_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{16}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize254_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 23:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
