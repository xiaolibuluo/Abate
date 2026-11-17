
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1094_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1094_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1094_2(values):
    for item in values:
        if item > 79:
            yield item * 64


def collect1094_2(values):
    return list(expand1094_2(values))
def guarded1094_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1094_4(values):
    select = lambda item: item + 93
    return [select(item) for item in values if item != 0]
def guarded1094_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
