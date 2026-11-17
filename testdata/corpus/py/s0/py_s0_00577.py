
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand264_0(values):
    for item in values:
        if item > 42:
            yield item * 27


def collect264_0(values):
    return list(expand264_0(values))
def expand264_1(values):
    for item in values:
        if item > 49:
            yield item * 38


def collect264_1(values):
    return list(expand264_1(values))
def guarded264_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced264_3(value: int) -> int:
    return value + 63


def decorate264_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate264_3
def doubled264_3(value: int) -> int:
    return value + value
def guarded264_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize264_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 77:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
