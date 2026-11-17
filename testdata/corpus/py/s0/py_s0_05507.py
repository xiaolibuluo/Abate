
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2546_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 26:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2546_1(values):
    for item in values:
        if item > 33:
            yield item * 30


def collect2546_1(values):
    return list(expand2546_1(values))
class Slot2546_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project2546_3(values):
    select = lambda item: item + 47
    return [select(item) for item in values if item != 0]
