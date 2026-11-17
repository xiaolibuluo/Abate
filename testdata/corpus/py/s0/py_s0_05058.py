
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2352_0(values):
    for item in values:
        if item > 26:
            yield item * 64


def collect2352_0(values):
    return list(expand2352_0(values))
class Slot2352_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{33}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2352_2(values):
    for item in values:
        if item > 40:
            yield item * 25


def collect2352_2(values):
    return list(expand2352_2(values))
class Slot2352_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{47}"

    @property
    def doubled(self) -> int:
        return self.size * 2
