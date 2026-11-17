
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize641_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 93:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot641_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{100}"

    @property
    def doubled(self) -> int:
        return self.size * 2
