
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot792_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{19}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project792_1(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
def summarize792_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 33:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair792_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 40}
