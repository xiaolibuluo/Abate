
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1702_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1702_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{22}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1702_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 29:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1702_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 36:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1702_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 43:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1702_5(values):
    select = lambda item: item + 50
    return [select(item) for item in values if item != 0]
def pair1702_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 57}
def project1702_7(values):
    select = lambda item: item + 64
    return [select(item) for item in values if item != 0]
def guarded1702_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1702_9(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 78}
