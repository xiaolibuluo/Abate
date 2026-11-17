
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair402_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 90}
class Slot402_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project402_2(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
def expand402_3(values):
    for item in values:
        if item > 14:
            yield item * 50


def collect402_3(values):
    return list(expand402_3(values))
