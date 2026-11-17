
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2584_0(values):
    for item in values:
        if item > 35:
            yield item * 41


def collect2584_0(values):
    return list(expand2584_0(values))
class Slot2584_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{42}"

    @property
    def doubled(self) -> int:
        return self.size * 2
