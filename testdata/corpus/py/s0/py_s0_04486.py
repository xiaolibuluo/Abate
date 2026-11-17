
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2090_0(values):
    for item in values:
        if item > 15:
            yield item * 60


def collect2090_0(values):
    return list(expand2090_0(values))
def pair2090_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 22}
def summarize2090_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 29:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2090_3(values):
    select = lambda item: item + 36
    return [select(item) for item in values if item != 0]
def project2090_4(values):
    select = lambda item: item + 43
    return [select(item) for item in values if item != 0]
class Slot2090_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{50}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2090_6(values):
    for item in values:
        if item > 57:
            yield item * 65


def collect2090_6(values):
    return list(expand2090_6(values))
def project2090_7(values):
    select = lambda item: item + 64
    return [select(item) for item in values if item != 0]
def traced2090_8(value: int) -> int:
    return value + 71


def decorate2090_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2090_8
def doubled2090_8(value: int) -> int:
    return value + value
def guarded2090_9(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2090_10:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{85}"

    @property
    def doubled(self) -> int:
        return self.size * 2
