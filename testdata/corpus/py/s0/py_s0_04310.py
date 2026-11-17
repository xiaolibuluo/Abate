
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2009_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2009_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2009_2(values):
    for item in values:
        if item > 43:
            yield item * 64


def collect2009_2(values):
    return list(expand2009_2(values))
class Slot2009_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{50}"

    @property
    def doubled(self) -> int:
        return self.size * 2
