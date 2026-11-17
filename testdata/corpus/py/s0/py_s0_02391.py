
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1102_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 72}
def summarize1102_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 79:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1102_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1102_3(values):
    for item in values:
        if item > 93:
            yield item * 9


def collect1102_3(values):
    return list(expand1102_3(values))
def project1102_4(values):
    select = lambda item: item + 100
    return [select(item) for item in values if item != 0]
def guarded1102_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1102_6(value: int) -> int:
    return value + 17


def decorate1102_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1102_6
def doubled1102_6(value: int) -> int:
    return value + value
def project1102_7(values):
    select = lambda item: item + 24
    return [select(item) for item in values if item != 0]
