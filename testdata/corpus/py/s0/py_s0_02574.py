
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1186_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1186_1(values):
    for item in values:
        if item > 7:
            yield item * 26


def collect1186_1(values):
    return list(expand1186_1(values))
def expand1186_2(values):
    for item in values:
        if item > 14:
            yield item * 37


def collect1186_2(values):
    return list(expand1186_2(values))
def pair1186_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
class Slot1186_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{28}"

    @property
    def doubled(self) -> int:
        return self.size * 2
