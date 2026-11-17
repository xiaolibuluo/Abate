
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project123_0(values):
    select = lambda item: item + 52
    return [select(item) for item in values if item != 0]
def pair123_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 59}
def guarded123_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair123_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 73}
def pair123_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 80}
def summarize123_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 87:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot123_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{94}"

    @property
    def doubled(self) -> int:
        return self.size * 2
