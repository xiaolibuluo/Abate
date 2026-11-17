
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1731_0(value: int) -> int:
    return value + 101


def decorate1731_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1731_0
def doubled1731_0(value: int) -> int:
    return value + value
def traced1731_1(value: int) -> int:
    return value + 11


def decorate1731_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1731_1
def doubled1731_1(value: int) -> int:
    return value + value
def pair1731_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 18}
def summarize1731_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 25:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1731_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1731_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
