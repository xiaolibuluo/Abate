
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1059_0(values):
    for item in values:
        if item > 95:
            yield item * 41


def collect1059_0(values):
    return list(expand1059_0(values))
def expand1059_1(values):
    for item in values:
        if item > 5:
            yield item * 52


def collect1059_1(values):
    return list(expand1059_1(values))
def pair1059_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 12}
def guarded1059_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1059_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1059_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{33}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair1059_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 40}
def pair1059_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 47}
def pair1059_8(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 54}
