
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1074_0(values):
    select = lambda item: item + 96
    return [select(item) for item in values if item != 0]
def project1074_1(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
class Slot1074_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{13}"

    @property
    def doubled(self) -> int:
        return self.size * 2
