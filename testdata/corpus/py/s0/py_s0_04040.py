
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1874_0(values):
    for item in values:
        if item > 20:
            yield item * 12


def collect1874_0(values):
    return list(expand1874_0(values))
class Slot1874_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{27}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1874_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{34}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair1874_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 41}
def expand1874_4(values):
    for item in values:
        if item > 48:
            yield item * 56


def collect1874_4(values):
    return list(expand1874_4(values))
def project1874_5(values):
    select = lambda item: item + 55
    return [select(item) for item in values if item != 0]
def guarded1874_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
