
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1902_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{93}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1902_1(values):
    for item in values:
        if item > 100:
            yield item * 36


def collect1902_1(values):
    return list(expand1902_1(values))
