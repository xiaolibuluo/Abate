
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2006_0(value: int) -> int:
    return value + 87


def decorate2006_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2006_0
def doubled2006_0(value: int) -> int:
    return value + value
class Slot2006_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{94}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2006_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
