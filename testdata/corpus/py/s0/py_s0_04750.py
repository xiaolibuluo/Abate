
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2202_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 16}
def guarded2202_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2202_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2202_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{37}"

    @property
    def doubled(self) -> int:
        return self.size * 2
