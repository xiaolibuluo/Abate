
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1778_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1778_1(value: int) -> int:
    return value + 40


def decorate1778_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1778_1
def doubled1778_1(value: int) -> int:
    return value + value
class Slot1778_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{47}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1778_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1778_4(values):
    select = lambda item: item + 61
    return [select(item) for item in values if item != 0]
def summarize1778_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 68:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
