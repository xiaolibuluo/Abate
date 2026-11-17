
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize331_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 40:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project331_1(values):
    select = lambda item: item + 47
    return [select(item) for item in values if item != 0]
def summarize331_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 54:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize331_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 61:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project331_4(values):
    select = lambda item: item + 68
    return [select(item) for item in values if item != 0]
def traced331_5(value: int) -> int:
    return value + 75


def decorate331_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate331_5
def doubled331_5(value: int) -> int:
    return value + value
def summarize331_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 82:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded331_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
