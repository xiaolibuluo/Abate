
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot549_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{61}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded549_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project549_2(values):
    select = lambda item: item + 75
    return [select(item) for item in values if item != 0]
def expand549_3(values):
    for item in values:
        if item > 82:
            yield item * 42


def collect549_3(values):
    return list(expand549_3(values))
