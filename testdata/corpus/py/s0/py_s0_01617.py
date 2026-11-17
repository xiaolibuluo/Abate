
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot736_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{67}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand736_1(values):
    for item in values:
        if item > 74:
            yield item * 48


def collect736_1(values):
    return list(expand736_1(values))
