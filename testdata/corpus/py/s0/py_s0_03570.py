
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1644_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{37}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1644_1(value: int) -> int:
    return value + 44


def decorate1644_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1644_1
def doubled1644_1(value: int) -> int:
    return value + value
def guarded1644_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1644_3(value: int) -> int:
    return value + 58


def decorate1644_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1644_3
def doubled1644_3(value: int) -> int:
    return value + value
def guarded1644_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1644_5(value: int) -> int:
    return value + 72


def decorate1644_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1644_5
def doubled1644_5(value: int) -> int:
    return value + value
def traced1644_6(value: int) -> int:
    return value + 79


def decorate1644_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1644_6
def doubled1644_6(value: int) -> int:
    return value + value
def guarded1644_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
