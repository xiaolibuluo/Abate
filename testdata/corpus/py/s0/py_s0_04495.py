
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2096_0(values):
    select = lambda item: item + 93
    return [select(item) for item in values if item != 0]
def pair2096_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 100}
class Slot2096_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{10}"

    @property
    def doubled(self) -> int:
        return self.size * 2
