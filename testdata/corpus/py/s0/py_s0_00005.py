
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot3_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{44}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project3_1(values):
    select = lambda item: item + 51
    return [select(item) for item in values if item != 0]
