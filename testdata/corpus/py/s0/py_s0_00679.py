
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot314_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{13}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize314_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair314_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 27}
def project314_3(values):
    select = lambda item: item + 34
    return [select(item) for item in values if item != 0]
def guarded314_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize314_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 48:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
