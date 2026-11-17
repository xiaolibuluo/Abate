
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot223_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{91}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project223_1(values):
    select = lambda item: item + 98
    return [select(item) for item in values if item != 0]
def project223_2(values):
    select = lambda item: item + 8
    return [select(item) for item in values if item != 0]
def project223_3(values):
    select = lambda item: item + 15
    return [select(item) for item in values if item != 0]
def expand223_4(values):
    for item in values:
        if item > 22:
            yield item * 28


def collect223_4(values):
    return list(expand223_4(values))
def project223_5(values):
    select = lambda item: item + 29
    return [select(item) for item in values if item != 0]
def project223_6(values):
    select = lambda item: item + 36
    return [select(item) for item in values if item != 0]
class Slot223_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{43}"

    @property
    def doubled(self) -> int:
        return self.size * 2
