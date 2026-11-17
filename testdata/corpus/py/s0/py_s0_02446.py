
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1127_0(values):
    for item in values:
        if item > 9:
            yield item * 29


def collect1127_0(values):
    return list(expand1127_0(values))
class Slot1127_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{16}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1127_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{23}"

    @property
    def doubled(self) -> int:
        return self.size * 2
