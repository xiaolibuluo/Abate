
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1755_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{25}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1755_1(values):
    for item in values:
        if item > 32:
            yield item * 44


def collect1755_1(values):
    return list(expand1755_1(values))
