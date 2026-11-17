
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1252_0(values):
    select = lambda item: item + 82
    return [select(item) for item in values if item != 0]
def guarded1252_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1252_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1252_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{6}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair1252_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
def guarded1252_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1252_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 27:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1252_7(value: int) -> int:
    return value + 34


def decorate1252_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1252_7
def doubled1252_7(value: int) -> int:
    return value + value
