
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced279_0(value: int) -> int:
    return value + 43


def decorate279_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate279_0
def doubled279_0(value: int) -> int:
    return value + value
def summarize279_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 50:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project279_2(values):
    select = lambda item: item + 57
    return [select(item) for item in values if item != 0]
def summarize279_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded279_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize279_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand279_6(values):
    for item in values:
        if item > 85:
            yield item * 15


def collect279_6(values):
    return list(expand279_6(values))
