
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1048_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 49}
def guarded1048_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1048_2(value: int) -> int:
    return value + 63


def decorate1048_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1048_2
def doubled1048_2(value: int) -> int:
    return value + value
def guarded1048_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1048_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1048_5(values):
    for item in values:
        if item > 84:
            yield item * 19


def collect1048_5(values):
    return list(expand1048_5(values))
def summarize1048_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 91:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
