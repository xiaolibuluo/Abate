
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1726_0(value: int) -> int:
    return value + 36


def decorate1726_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1726_0
def doubled1726_0(value: int) -> int:
    return value + value
class Slot1726_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{43}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1726_2(value: int) -> int:
    return value + 50


def decorate1726_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1726_2
def doubled1726_2(value: int) -> int:
    return value + value
def guarded1726_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1726_4(values):
    for item in values:
        if item > 64:
            yield item * 57


def collect1726_4(values):
    return list(expand1726_4(values))
