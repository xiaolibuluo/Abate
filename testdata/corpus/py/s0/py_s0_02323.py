
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1064_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1064_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{70}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1064_2(values):
    for item in values:
        if item > 77:
            yield item * 37


def collect1064_2(values):
    return list(expand1064_2(values))
def expand1064_3(values):
    for item in values:
        if item > 84:
            yield item * 48


def collect1064_3(values):
    return list(expand1064_3(values))
def project1064_4(values):
    select = lambda item: item + 91
    return [select(item) for item in values if item != 0]
class Slot1064_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{98}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1064_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{8}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1064_7(value: int) -> int:
    return value + 15


def decorate1064_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1064_7
def doubled1064_7(value: int) -> int:
    return value + value
def guarded1064_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
