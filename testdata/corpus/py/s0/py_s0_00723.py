
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded330_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced330_1(value: int) -> int:
    return value + 34


def decorate330_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate330_1
def doubled330_1(value: int) -> int:
    return value + value
class Slot330_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{41}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand330_3(values):
    for item in values:
        if item > 48:
            yield item * 34


def collect330_3(values):
    return list(expand330_3(values))
def guarded330_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced330_5(value: int) -> int:
    return value + 62


def decorate330_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate330_5
def doubled330_5(value: int) -> int:
    return value + value
def guarded330_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot330_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{76}"

    @property
    def doubled(self) -> int:
        return self.size * 2
