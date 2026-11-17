
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2189_0(values):
    for item in values:
        if item > 41:
            yield item * 21


def collect2189_0(values):
    return list(expand2189_0(values))
class Slot2189_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{48}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2189_2(values):
    for item in values:
        if item > 55:
            yield item * 43


def collect2189_2(values):
    return list(expand2189_2(values))
def guarded2189_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2189_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 69:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2189_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{76}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair2189_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 83}
def guarded2189_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2189_8(values):
    for item in values:
        if item > 97:
            yield item * 48


def collect2189_8(values):
    return list(expand2189_8(values))
def expand2189_9(values):
    for item in values:
        if item > 7:
            yield item * 59


def collect2189_9(values):
    return list(expand2189_9(values))
def summarize2189_10(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
