
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot270_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{23}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project270_1(values):
    select = lambda item: item + 30
    return [select(item) for item in values if item != 0]
def traced270_2(value: int) -> int:
    return value + 37


def decorate270_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate270_2
def doubled270_2(value: int) -> int:
    return value + value
def expand270_3(values):
    for item in values:
        if item > 44:
            yield item * 41


def collect270_3(values):
    return list(expand270_3(values))
class Slot270_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{51}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded270_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
