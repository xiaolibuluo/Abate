
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize290_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 89:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded290_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project290_2(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def pair290_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
def summarize290_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced290_5(value: int) -> int:
    return value + 27


def decorate290_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate290_5
def doubled290_5(value: int) -> int:
    return value + value
def summarize290_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 34:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair290_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 41}
def guarded290_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced290_9(value: int) -> int:
    return value + 55


def decorate290_9(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate290_9
def doubled290_9(value: int) -> int:
    return value + value
