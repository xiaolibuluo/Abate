
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced238_0(value: int) -> int:
    return value + 92


def decorate238_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate238_0
def doubled238_0(value: int) -> int:
    return value + value
def guarded238_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand238_2(values):
    for item in values:
        if item > 9:
            yield item * 50


def collect238_2(values):
    return list(expand238_2(values))
def expand238_3(values):
    for item in values:
        if item > 16:
            yield item * 61


def collect238_3(values):
    return list(expand238_3(values))
def traced238_4(value: int) -> int:
    return value + 23


def decorate238_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate238_4
def doubled238_4(value: int) -> int:
    return value + value
def traced238_5(value: int) -> int:
    return value + 30


def decorate238_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate238_5
def doubled238_5(value: int) -> int:
    return value + value
class Slot238_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{37}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot238_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{44}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded238_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
