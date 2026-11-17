
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1671_0(values):
    select = lambda item: item + 97
    return [select(item) for item in values if item != 0]
def summarize1671_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 7:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1671_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{14}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1671_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
