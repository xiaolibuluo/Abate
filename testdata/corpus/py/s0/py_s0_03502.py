
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1616_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1616_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{68}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1616_2(values):
    for item in values:
        if item > 75:
            yield item * 58


def collect1616_2(values):
    return list(expand1616_2(values))
