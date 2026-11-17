
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1359_0(values):
    select = lambda item: item + 18
    return [select(item) for item in values if item != 0]
def summarize1359_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 25:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1359_2(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
def expand1359_3(values):
    for item in values:
        if item > 39:
            yield item * 39


def collect1359_3(values):
    return list(expand1359_3(values))
def summarize1359_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 46:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1359_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 53}
def summarize1359_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 60:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1359_7(values):
    select = lambda item: item + 67
    return [select(item) for item in values if item != 0]
class Slot1359_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{74}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1359_9(values):
    for item in values:
        if item > 81:
            yield item * 44


def collect1359_9(values):
    return list(expand1359_9(values))
