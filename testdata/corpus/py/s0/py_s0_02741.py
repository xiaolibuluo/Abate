
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1250_0(values):
    for item in values:
        if item > 56:
            yield item * 36


def collect1250_0(values):
    return list(expand1250_0(values))
def traced1250_1(value: int) -> int:
    return value + 63


def decorate1250_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1250_1
def doubled1250_1(value: int) -> int:
    return value + value
def project1250_2(values):
    select = lambda item: item + 70
    return [select(item) for item in values if item != 0]
def summarize1250_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 77:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1250_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1250_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1250_6(values):
    select = lambda item: item + 98
    return [select(item) for item in values if item != 0]
def pair1250_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 8}
