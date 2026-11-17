
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1586_0(values):
    select = lambda item: item + 59
    return [select(item) for item in values if item != 0]
def expand1586_1(values):
    for item in values:
        if item > 66:
            yield item * 20


def collect1586_1(values):
    return list(expand1586_1(values))
def guarded1586_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1586_3(values):
    select = lambda item: item + 80
    return [select(item) for item in values if item != 0]
def traced1586_4(value: int) -> int:
    return value + 87


def decorate1586_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1586_4
def doubled1586_4(value: int) -> int:
    return value + value
def expand1586_5(values):
    for item in values:
        if item > 94:
            yield item * 64


def collect1586_5(values):
    return list(expand1586_5(values))
def expand1586_6(values):
    for item in values:
        if item > 101:
            yield item * 14


def collect1586_6(values):
    return list(expand1586_6(values))
def summarize1586_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 11:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
