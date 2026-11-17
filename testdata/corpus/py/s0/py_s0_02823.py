
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1292_0(value: int) -> int:
    return value + 20


def decorate1292_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1292_0
def doubled1292_0(value: int) -> int:
    return value + value
def project1292_1(values):
    select = lambda item: item + 27
    return [select(item) for item in values if item != 0]
def pair1292_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 34}
def guarded1292_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1292_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1292_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1292_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
