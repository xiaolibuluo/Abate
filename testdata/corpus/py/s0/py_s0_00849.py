
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair387_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 89}
def expand387_1(values):
    for item in values:
        if item > 96:
            yield item * 45


def collect387_1(values):
    return list(expand387_1(values))
def project387_2(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def summarize387_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand387_4(values):
    for item in values:
        if item > 20:
            yield item * 17


def collect387_4(values):
    return list(expand387_4(values))
class Slot387_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{27}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand387_6(values):
    for item in values:
        if item > 34:
            yield item * 39


def collect387_6(values):
    return list(expand387_6(values))
def guarded387_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair387_8(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 48}
def guarded387_9(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot387_10:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{62}"

    @property
    def doubled(self) -> int:
        return self.size * 2
