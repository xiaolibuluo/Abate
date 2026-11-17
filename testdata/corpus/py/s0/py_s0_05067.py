
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2356_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2356_1(value: int) -> int:
    return value + 85


def decorate2356_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2356_1
def doubled2356_1(value: int) -> int:
    return value + value
class Slot2356_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{92}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project2356_3(values):
    select = lambda item: item + 99
    return [select(item) for item in values if item != 0]
def project2356_4(values):
    select = lambda item: item + 9
    return [select(item) for item in values if item != 0]
class Slot2356_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{16}"

    @property
    def doubled(self) -> int:
        return self.size * 2
