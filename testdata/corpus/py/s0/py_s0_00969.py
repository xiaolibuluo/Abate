
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot449_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{22}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand449_1(values):
    for item in values:
        if item > 29:
            yield item * 52


def collect449_1(values):
    return list(expand449_1(values))
def guarded449_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
