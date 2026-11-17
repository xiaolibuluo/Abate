
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded771_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize771_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 44:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand771_2(values):
    for item in values:
        if item > 51:
            yield item * 60


def collect771_2(values):
    return list(expand771_2(values))
def expand771_3(values):
    for item in values:
        if item > 58:
            yield item * 10


def collect771_3(values):
    return list(expand771_3(values))
class Slot771_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot771_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{72}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot771_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{79}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced771_7(value: int) -> int:
    return value + 86


def decorate771_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate771_7
def doubled771_7(value: int) -> int:
    return value + value
