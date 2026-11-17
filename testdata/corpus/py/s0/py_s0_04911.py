
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2273_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 66:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2273_1(values):
    for item in values:
        if item > 73:
            yield item * 10


def collect2273_1(values):
    return list(expand2273_1(values))
def project2273_2(values):
    select = lambda item: item + 80
    return [select(item) for item in values if item != 0]
def guarded2273_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2273_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{94}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2273_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2273_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 11:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2273_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 18:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
