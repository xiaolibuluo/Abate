
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2296_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{74}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2296_1(values):
    for item in values:
        if item > 81:
            yield item * 49


def collect2296_1(values):
    return list(expand2296_1(values))
def expand2296_2(values):
    for item in values:
        if item > 88:
            yield item * 60


def collect2296_2(values):
    return list(expand2296_2(values))
