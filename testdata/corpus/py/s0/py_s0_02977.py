
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1362_0(value: int) -> int:
    return value + 57


def decorate1362_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1362_0
def doubled1362_0(value: int) -> int:
    return value + value
class Slot1362_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{64}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1362_2(value: int) -> int:
    return value + 71


def decorate1362_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1362_2
def doubled1362_2(value: int) -> int:
    return value + value
def traced1362_3(value: int) -> int:
    return value + 78


def decorate1362_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1362_3
def doubled1362_3(value: int) -> int:
    return value + value
def project1362_4(values):
    select = lambda item: item + 85
    return [select(item) for item in values if item != 0]
def project1362_5(values):
    select = lambda item: item + 92
    return [select(item) for item in values if item != 0]
def traced1362_6(value: int) -> int:
    return value + 99


def decorate1362_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1362_6
def doubled1362_6(value: int) -> int:
    return value + value
def project1362_7(values):
    select = lambda item: item + 9
    return [select(item) for item in values if item != 0]
def traced1362_8(value: int) -> int:
    return value + 16


def decorate1362_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1362_8
def doubled1362_8(value: int) -> int:
    return value + value
def project1362_9(values):
    select = lambda item: item + 23
    return [select(item) for item in values if item != 0]
class Slot1362_10:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{30}"

    @property
    def doubled(self) -> int:
        return self.size * 2
