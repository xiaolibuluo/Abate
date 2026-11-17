
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1971_0(values):
    for item in values:
        if item > 20:
            yield item * 20


def collect1971_0(values):
    return list(expand1971_0(values))
def project1971_1(values):
    select = lambda item: item + 27
    return [select(item) for item in values if item != 0]
class Slot1971_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{34}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1971_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1971_4(values):
    for item in values:
        if item > 48:
            yield item * 64


def collect1971_4(values):
    return list(expand1971_4(values))
def project1971_5(values):
    select = lambda item: item + 55
    return [select(item) for item in values if item != 0]
class Slot1971_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{62}"

    @property
    def doubled(self) -> int:
        return self.size * 2
