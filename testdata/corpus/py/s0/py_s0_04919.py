
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2279_0(values):
    for item in values:
        if item > 47:
            yield item * 41


def collect2279_0(values):
    return list(expand2279_0(values))
def guarded2279_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2279_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{61}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project2279_3(values):
    select = lambda item: item + 68
    return [select(item) for item in values if item != 0]
def guarded2279_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2279_5(value: int) -> int:
    return value + 82


def decorate2279_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2279_5
def doubled2279_5(value: int) -> int:
    return value + value
def guarded2279_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2279_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 96:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2279_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{6}"

    @property
    def doubled(self) -> int:
        return self.size * 2
