
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1466_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{51}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project1466_1(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
def expand1466_2(values):
    for item in values:
        if item > 65:
            yield item * 45


def collect1466_2(values):
    return list(expand1466_2(values))
