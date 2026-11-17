
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2139_0(values):
    for item in values:
        if item > 70:
            yield item * 37


def collect2139_0(values):
    return list(expand2139_0(values))
def summarize2139_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 77:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2139_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 84:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2139_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{91}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2139_4(values):
    for item in values:
        if item > 98:
            yield item * 20


def collect2139_4(values):
    return list(expand2139_4(values))
def pair2139_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 8}
def summarize2139_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 15:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2139_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2139_8(values):
    for item in values:
        if item > 29:
            yield item * 64


def collect2139_8(values):
    return list(expand2139_8(values))
class Slot2139_9:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{36}"

    @property
    def doubled(self) -> int:
        return self.size * 2
