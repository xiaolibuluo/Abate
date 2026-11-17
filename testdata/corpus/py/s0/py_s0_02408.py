
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1111_0(values):
    for item in values:
        if item > 92:
            yield item * 39


def collect1111_0(values):
    return list(expand1111_0(values))
def guarded1111_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1111_2(value: int) -> int:
    return value + 9


def decorate1111_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1111_2
def doubled1111_2(value: int) -> int:
    return value + value
def expand1111_3(values):
    for item in values:
        if item > 16:
            yield item * 11


def collect1111_3(values):
    return list(expand1111_3(values))
def guarded1111_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1111_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 30}
def summarize1111_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 37:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1111_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 44:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1111_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
