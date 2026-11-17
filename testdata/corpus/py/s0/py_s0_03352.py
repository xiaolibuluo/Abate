
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1546_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1546_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 31:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1546_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{38}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project1546_3(values):
    select = lambda item: item + 45
    return [select(item) for item in values if item != 0]
def summarize1546_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 52:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1546_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{59}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1546_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1546_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 73}
