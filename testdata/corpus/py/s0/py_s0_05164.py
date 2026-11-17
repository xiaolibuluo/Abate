
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2407_0(values):
    for item in values:
        if item > 62:
            yield item * 22


def collect2407_0(values):
    return list(expand2407_0(values))
def expand2407_1(values):
    for item in values:
        if item > 69:
            yield item * 33


def collect2407_1(values):
    return list(expand2407_1(values))
def pair2407_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 76}
def summarize2407_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 83:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2407_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{90}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2407_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
