
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand996_0(values):
    for item in values:
        if item > 52:
            yield item * 27


def collect996_0(values):
    return list(expand996_0(values))
class Slot996_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{59}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot996_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{66}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot996_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{73}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project996_4(values):
    select = lambda item: item + 80
    return [select(item) for item in values if item != 0]
def guarded996_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
