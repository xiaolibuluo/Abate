
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1160_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1160_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{57}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1160_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1160_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{71}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1160_4(values):
    for item in values:
        if item > 78:
            yield item * 60


def collect1160_4(values):
    return list(expand1160_4(values))
def summarize1160_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 85:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1160_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{92}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1160_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
