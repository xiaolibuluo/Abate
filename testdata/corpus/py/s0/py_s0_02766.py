
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1262_0(value: int) -> int:
    return value + 18


def decorate1262_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1262_0
def doubled1262_0(value: int) -> int:
    return value + value
def traced1262_1(value: int) -> int:
    return value + 25


def decorate1262_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1262_1
def doubled1262_1(value: int) -> int:
    return value + value
def expand1262_2(values):
    for item in values:
        if item > 32:
            yield item * 20


def collect1262_2(values):
    return list(expand1262_2(values))
def summarize1262_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 39:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1262_4(values):
    select = lambda item: item + 46
    return [select(item) for item in values if item != 0]
def project1262_5(values):
    select = lambda item: item + 53
    return [select(item) for item in values if item != 0]
def summarize1262_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 60:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1262_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 67:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1262_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
