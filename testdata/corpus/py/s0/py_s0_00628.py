
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded286_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded286_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot286_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{51}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot286_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{58}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot286_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project286_5(values):
    select = lambda item: item + 72
    return [select(item) for item in values if item != 0]
def traced286_6(value: int) -> int:
    return value + 79


def decorate286_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate286_6
def doubled286_6(value: int) -> int:
    return value + value
class Slot286_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{86}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced286_8(value: int) -> int:
    return value + 93


def decorate286_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate286_8
def doubled286_8(value: int) -> int:
    return value + value
