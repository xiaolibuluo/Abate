
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1389_0(values):
    select = lambda item: item + 20
    return [select(item) for item in values if item != 0]
def pair1389_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 27}
def expand1389_2(values):
    for item in values:
        if item > 34:
            yield item * 55


def collect1389_2(values):
    return list(expand1389_2(values))
def summarize1389_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 41:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1389_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{48}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1389_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1389_6(values):
    select = lambda item: item + 62
    return [select(item) for item in values if item != 0]
