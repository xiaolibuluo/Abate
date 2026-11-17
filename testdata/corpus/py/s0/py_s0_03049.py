
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1401_0(values):
    for item in values:
        if item > 79:
            yield item * 56


def collect1401_0(values):
    return list(expand1401_0(values))
def guarded1401_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1401_2(value: int) -> int:
    return value + 93


def decorate1401_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1401_2
def doubled1401_2(value: int) -> int:
    return value + value
def summarize1401_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1401_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{10}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1401_5(value: int) -> int:
    return value + 17


def decorate1401_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1401_5
def doubled1401_5(value: int) -> int:
    return value + value
def traced1401_6(value: int) -> int:
    return value + 24


def decorate1401_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1401_6
def doubled1401_6(value: int) -> int:
    return value + value
def project1401_7(values):
    select = lambda item: item + 31
    return [select(item) for item in values if item != 0]
