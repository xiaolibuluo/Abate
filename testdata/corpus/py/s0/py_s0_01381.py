
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize629_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 34:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot629_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{41}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded629_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize629_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand629_4(values):
    for item in values:
        if item > 62:
            yield item * 64


def collect629_4(values):
    return list(expand629_4(values))
def expand629_5(values):
    for item in values:
        if item > 69:
            yield item * 14


def collect629_5(values):
    return list(expand629_5(values))
def expand629_6(values):
    for item in values:
        if item > 76:
            yield item * 25


def collect629_6(values):
    return list(expand629_6(values))
def summarize629_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 83:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot629_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{90}"

    @property
    def doubled(self) -> int:
        return self.size * 2
