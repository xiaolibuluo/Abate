
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1090_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1090_1(value: int) -> int:
    return value + 20


def decorate1090_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1090_1
def doubled1090_1(value: int) -> int:
    return value + value
class Slot1090_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{27}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1090_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{34}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1090_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
