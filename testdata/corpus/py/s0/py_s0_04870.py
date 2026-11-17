
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2254_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2254_1(values):
    for item in values:
        if item > 20:
            yield item * 60


def collect2254_1(values):
    return list(expand2254_1(values))
def traced2254_2(value: int) -> int:
    return value + 27


def decorate2254_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2254_2
def doubled2254_2(value: int) -> int:
    return value + value
class Slot2254_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{34}"

    @property
    def doubled(self) -> int:
        return self.size * 2
