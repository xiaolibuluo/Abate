
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1626_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{94}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1626_1(values):
    for item in values:
        if item > 101:
            yield item * 56


def collect1626_1(values):
    return list(expand1626_1(values))
def expand1626_2(values):
    for item in values:
        if item > 11:
            yield item * 67


def collect1626_2(values):
    return list(expand1626_2(values))
def expand1626_3(values):
    for item in values:
        if item > 18:
            yield item * 17


def collect1626_3(values):
    return list(expand1626_3(values))
class Slot1626_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{25}"

    @property
    def doubled(self) -> int:
        return self.size * 2
