
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2442_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{32}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2442_1(values):
    for item in values:
        if item > 39:
            yield item * 34


def collect2442_1(values):
    return list(expand2442_1(values))
class Slot2442_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{46}"

    @property
    def doubled(self) -> int:
        return self.size * 2
