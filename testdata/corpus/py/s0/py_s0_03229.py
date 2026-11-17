
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1492_0(values):
    for item in values:
        if item > 98:
            yield item * 22


def collect1492_0(values):
    return list(expand1492_0(values))
def guarded1492_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1492_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1492_3(value: int) -> int:
    return value + 22


def decorate1492_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1492_3
def doubled1492_3(value: int) -> int:
    return value + value
def summarize1492_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 29:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1492_5(value: int) -> int:
    return value + 36


def decorate1492_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1492_5
def doubled1492_5(value: int) -> int:
    return value + value
def pair1492_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 43}
