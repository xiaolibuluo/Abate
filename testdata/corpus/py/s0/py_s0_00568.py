
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project260_0(values):
    select = lambda item: item + 87
    return [select(item) for item in values if item != 0]
def traced260_1(value: int) -> int:
    return value + 94


def decorate260_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate260_1
def doubled260_1(value: int) -> int:
    return value + value
def traced260_2(value: int) -> int:
    return value + 101


def decorate260_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate260_2
def doubled260_2(value: int) -> int:
    return value + value
class Slot260_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{11}"

    @property
    def doubled(self) -> int:
        return self.size * 2
