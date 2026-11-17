
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2492_0(values):
    for item in values:
        if item > 100:
            yield item * 68


def collect2492_0(values):
    return list(expand2492_0(values))
def project2492_1(values):
    select = lambda item: item + 10
    return [select(item) for item in values if item != 0]
class Slot2492_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{17}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project2492_3(values):
    select = lambda item: item + 24
    return [select(item) for item in values if item != 0]
