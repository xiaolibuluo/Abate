
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2128_0(value: int) -> int:
    return value + 24


def decorate2128_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2128_0
def doubled2128_0(value: int) -> int:
    return value + value
class Slot2128_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{31}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2128_2(value: int) -> int:
    return value + 38


def decorate2128_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2128_2
def doubled2128_2(value: int) -> int:
    return value + value
def guarded2128_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2128_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{52}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2128_5(value: int) -> int:
    return value + 59


def decorate2128_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2128_5
def doubled2128_5(value: int) -> int:
    return value + value
def project2128_6(values):
    select = lambda item: item + 66
    return [select(item) for item in values if item != 0]
def summarize2128_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 73:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
