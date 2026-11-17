
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1527_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1527_1(values):
    select = lambda item: item + 75
    return [select(item) for item in values if item != 0]
def guarded1527_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1527_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{89}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1527_4(values):
    for item in values:
        if item > 96:
            yield item * 67


def collect1527_4(values):
    return list(expand1527_4(values))
