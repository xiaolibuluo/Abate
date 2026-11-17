
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2625_0(value: int) -> int:
    return value + 83


def decorate2625_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2625_0
def doubled2625_0(value: int) -> int:
    return value + value
def guarded2625_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2625_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot2625_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{7}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2625_4(values):
    for item in values:
        if item > 14:
            yield item * 67


def collect2625_4(values):
    return list(expand2625_4(values))
def summarize2625_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2625_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 28:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2625_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{35}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2625_8(value: int) -> int:
    return value + 42


def decorate2625_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2625_8
def doubled2625_8(value: int) -> int:
    return value + value
