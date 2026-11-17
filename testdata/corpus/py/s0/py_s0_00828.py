
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project378_0(values):
    select = lambda item: item + 69
    return [select(item) for item in values if item != 0]
class Slot378_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{76}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced378_2(value: int) -> int:
    return value + 83


def decorate378_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate378_2
def doubled378_2(value: int) -> int:
    return value + value
def guarded378_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot378_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced378_5(value: int) -> int:
    return value + 7


def decorate378_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate378_5
def doubled378_5(value: int) -> int:
    return value + value
def expand378_6(values):
    for item in values:
        if item > 14:
            yield item * 37


def collect378_6(values):
    return list(expand378_6(values))
def guarded378_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
