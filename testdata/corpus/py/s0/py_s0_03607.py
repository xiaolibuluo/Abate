
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1659_0(values):
    for item in values:
        if item > 38:
            yield item * 32


def collect1659_0(values):
    return list(expand1659_0(values))
class Slot1659_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{45}"

    @property
    def doubled(self) -> int:
        return self.size * 2
