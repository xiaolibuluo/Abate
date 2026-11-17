
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand15_0(values):
    for item in values:
        if item > 6:
            yield item * 53


def collect15_0(values):
    return list(expand15_0(values))
class Slot15_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{13}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced15_2(value: int) -> int:
    return value + 20


def decorate15_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate15_2
def doubled15_2(value: int) -> int:
    return value + value
def guarded15_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
