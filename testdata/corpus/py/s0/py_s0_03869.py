
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1788_0(values):
    for item in values:
        if item > 66:
            yield item * 20


def collect1788_0(values):
    return list(expand1788_0(values))
def pair1788_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 73}
class Slot1788_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{80}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1788_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 87:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1788_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1788_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{101}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1788_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1788_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 18:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
