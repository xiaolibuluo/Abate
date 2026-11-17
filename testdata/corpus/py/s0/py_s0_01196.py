
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot552_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{100}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded552_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot552_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{17}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand552_3(values):
    for item in values:
        if item > 24:
            yield item * 63


def collect552_3(values):
    return list(expand552_3(values))
