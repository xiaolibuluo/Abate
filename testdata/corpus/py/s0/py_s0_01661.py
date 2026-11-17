
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot756_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{36}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot756_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{43}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot756_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{50}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project756_3(values):
    select = lambda item: item + 57
    return [select(item) for item in values if item != 0]
def project756_4(values):
    select = lambda item: item + 64
    return [select(item) for item in values if item != 0]
def project756_5(values):
    select = lambda item: item + 71
    return [select(item) for item in values if item != 0]
class Slot756_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{78}"

    @property
    def doubled(self) -> int:
        return self.size * 2
