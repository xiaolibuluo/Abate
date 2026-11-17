
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1420_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1420_1(values):
    select = lambda item: item + 42
    return [select(item) for item in values if item != 0]
def project1420_2(values):
    select = lambda item: item + 49
    return [select(item) for item in values if item != 0]
class Slot1420_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{56}"

    @property
    def doubled(self) -> int:
        return self.size * 2
