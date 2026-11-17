
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded580_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced580_1(value: int) -> int:
    return value + 83


def decorate580_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate580_1
def doubled580_1(value: int) -> int:
    return value + value
def summarize580_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 90:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot580_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project580_4(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
def summarize580_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded580_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded580_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded580_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
