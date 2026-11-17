
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1301_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1301_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{47}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1301_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 54:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1301_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 61:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1301_4(values):
    select = lambda item: item + 68
    return [select(item) for item in values if item != 0]
def traced1301_5(value: int) -> int:
    return value + 75


def decorate1301_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1301_5
def doubled1301_5(value: int) -> int:
    return value + value
def expand1301_6(values):
    for item in values:
        if item > 82:
            yield item * 32


def collect1301_6(values):
    return list(expand1301_6(values))
def summarize1301_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 89:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
