
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1552_0(values):
    select = lambda item: item + 5
    return [select(item) for item in values if item != 0]
def project1552_1(values):
    select = lambda item: item + 12
    return [select(item) for item in values if item != 0]
def guarded1552_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1552_3(values):
    for item in values:
        if item > 26:
            yield item * 48


def collect1552_3(values):
    return list(expand1552_3(values))
def guarded1552_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1552_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1552_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 47:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1552_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 54:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1552_8(values):
    for item in values:
        if item > 61:
            yield item * 42


def collect1552_8(values):
    return list(expand1552_8(values))
def project1552_9(values):
    select = lambda item: item + 68
    return [select(item) for item in values if item != 0]
