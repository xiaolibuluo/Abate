
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2297_0(values):
    select = lambda item: item + 87
    return [select(item) for item in values if item != 0]
class Slot2297_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{94}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2297_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2297_3(values):
    select = lambda item: item + 11
    return [select(item) for item in values if item != 0]
