
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2635_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{19}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2635_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2635_2(value: int) -> int:
    return value + 33


def decorate2635_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2635_2
def doubled2635_2(value: int) -> int:
    return value + value
