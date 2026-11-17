
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1000_0(value: int) -> int:
    return value + 7


def decorate1000_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1000_0
def doubled1000_0(value: int) -> int:
    return value + value
def traced1000_1(value: int) -> int:
    return value + 14


def decorate1000_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1000_1
def doubled1000_1(value: int) -> int:
    return value + value
def guarded1000_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1000_3(values):
    select = lambda item: item + 28
    return [select(item) for item in values if item != 0]
def guarded1000_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1000_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{42}"

    @property
    def doubled(self) -> int:
        return self.size * 2
