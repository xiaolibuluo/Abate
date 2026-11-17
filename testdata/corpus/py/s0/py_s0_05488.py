
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2535_0(values):
    select = lambda item: item + 77
    return [select(item) for item in values if item != 0]
class Slot2535_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{84}"

    @property
    def doubled(self) -> int:
        return self.size * 2
