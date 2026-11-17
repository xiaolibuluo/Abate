
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand321_0(values):
    for item in values:
        if item > 7:
            yield item * 60


def collect321_0(values):
    return list(expand321_0(values))
def guarded321_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize321_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced321_3(value: int) -> int:
    return value + 28


def decorate321_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate321_3
def doubled321_3(value: int) -> int:
    return value + value
def expand321_4(values):
    for item in values:
        if item > 35:
            yield item * 43


def collect321_4(values):
    return list(expand321_4(values))
def traced321_5(value: int) -> int:
    return value + 42


def decorate321_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate321_5
def doubled321_5(value: int) -> int:
    return value + value
