
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair40_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 40}
class Slot40_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{47}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize40_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 54:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot40_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{61}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded40_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand40_5(values):
    for item in values:
        if item > 75:
            yield item * 39


def collect40_5(values):
    return list(expand40_5(values))
def summarize40_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 82:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded40_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
