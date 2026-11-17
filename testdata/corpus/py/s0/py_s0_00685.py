
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded317_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand317_1(values):
    for item in values:
        if item > 59:
            yield item * 43


def collect317_1(values):
    return list(expand317_1(values))
def expand317_2(values):
    for item in values:
        if item > 66:
            yield item * 54


def collect317_2(values):
    return list(expand317_2(values))
def summarize317_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 73:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand317_4(values):
    for item in values:
        if item > 80:
            yield item * 15


def collect317_4(values):
    return list(expand317_4(values))
def traced317_5(value: int) -> int:
    return value + 87


def decorate317_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate317_5
def doubled317_5(value: int) -> int:
    return value + value
