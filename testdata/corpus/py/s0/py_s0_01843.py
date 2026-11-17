
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand843_0(values):
    for item in values:
        if item > 100:
            yield item * 54


def collect843_0(values):
    return list(expand843_0(values))
def project843_1(values):
    select = lambda item: item + 10
    return [select(item) for item in values if item != 0]
class Slot843_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{17}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot843_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{24}"

    @property
    def doubled(self) -> int:
        return self.size * 2
