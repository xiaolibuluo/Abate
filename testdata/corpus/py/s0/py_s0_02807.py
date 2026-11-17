
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1284_0(value: int) -> int:
    return value + 13


def decorate1284_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1284_0
def doubled1284_0(value: int) -> int:
    return value + value
class Slot1284_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{20}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1284_2(values):
    for item in values:
        if item > 27:
            yield item * 52


def collect1284_2(values):
    return list(expand1284_2(values))
def guarded1284_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1284_4(values):
    select = lambda item: item + 41
    return [select(item) for item in values if item != 0]
