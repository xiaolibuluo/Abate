
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1066_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1066_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{96}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1066_2(values):
    for item in values:
        if item > 6:
            yield item * 51


def collect1066_2(values):
    return list(expand1066_2(values))
def expand1066_3(values):
    for item in values:
        if item > 13:
            yield item * 62


def collect1066_3(values):
    return list(expand1066_3(values))
def guarded1066_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1066_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 27:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1066_6(value: int) -> int:
    return value + 34


def decorate1066_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1066_6
def doubled1066_6(value: int) -> int:
    return value + value
def expand1066_7(values):
    for item in values:
        if item > 41:
            yield item * 45


def collect1066_7(values):
    return list(expand1066_7(values))
def traced1066_8(value: int) -> int:
    return value + 48


def decorate1066_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1066_8
def doubled1066_8(value: int) -> int:
    return value + value
