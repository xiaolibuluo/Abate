
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot189_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{37}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded189_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair189_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 51}
def guarded189_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded189_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded189_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand189_6(values):
    for item in values:
        if item > 79:
            yield item * 56


def collect189_6(values):
    return list(expand189_6(values))
def summarize189_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 86:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded189_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project189_9(values):
    select = lambda item: item + 100
    return [select(item) for item in values if item != 0]
