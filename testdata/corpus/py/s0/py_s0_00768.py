
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced353_0(value: int) -> int:
    return value + 35


def decorate353_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate353_0
def doubled353_0(value: int) -> int:
    return value + value
class Slot353_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{42}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded353_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
