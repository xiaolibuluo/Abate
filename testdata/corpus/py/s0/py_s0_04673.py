
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2165_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 20}
def guarded2165_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2165_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 34:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2165_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 41:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2165_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{48}"

    @property
    def doubled(self) -> int:
        return self.size * 2
