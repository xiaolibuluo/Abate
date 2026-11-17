
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand86_0(values):
    for item in values:
        if item > 56:
            yield item * 62


def collect86_0(values):
    return list(expand86_0(values))
def expand86_1(values):
    for item in values:
        if item > 63:
            yield item * 12


def collect86_1(values):
    return list(expand86_1(values))
def guarded86_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded86_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand86_4(values):
    for item in values:
        if item > 84:
            yield item * 45


def collect86_4(values):
    return list(expand86_4(values))
def guarded86_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot86_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{98}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project86_7(values):
    select = lambda item: item + 8
    return [select(item) for item in values if item != 0]
def pair86_8(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 15}
def project86_9(values):
    select = lambda item: item + 22
    return [select(item) for item in values if item != 0]
