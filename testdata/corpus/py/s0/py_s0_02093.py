
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project956_0(values):
    select = lambda item: item + 17
    return [select(item) for item in values if item != 0]
def summarize956_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 24:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced956_2(value: int) -> int:
    return value + 31


def decorate956_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate956_2
def doubled956_2(value: int) -> int:
    return value + value
def traced956_3(value: int) -> int:
    return value + 38


def decorate956_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate956_3
def doubled956_3(value: int) -> int:
    return value + value
class Slot956_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{45}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot956_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{52}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize956_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 59:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
