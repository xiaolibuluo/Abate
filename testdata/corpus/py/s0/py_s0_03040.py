
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1396_0(values):
    for item in values:
        if item > 14:
            yield item * 21


def collect1396_0(values):
    return list(expand1396_0(values))
def guarded1396_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1396_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 28:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1396_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 35:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1396_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1396_5(value: int) -> int:
    return value + 49


def decorate1396_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1396_5
def doubled1396_5(value: int) -> int:
    return value + value
def project1396_6(values):
    select = lambda item: item + 56
    return [select(item) for item in values if item != 0]
