
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1286_0(value: int) -> int:
    return value + 39


def decorate1286_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1286_0
def doubled1286_0(value: int) -> int:
    return value + value
def expand1286_1(values):
    for item in values:
        if item > 46:
            yield item * 55


def collect1286_1(values):
    return list(expand1286_1(values))
def expand1286_2(values):
    for item in values:
        if item > 53:
            yield item * 66


def collect1286_2(values):
    return list(expand1286_2(values))
class Slot1286_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{60}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1286_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
