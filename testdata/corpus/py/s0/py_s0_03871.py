
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1790_0(value: int) -> int:
    return value + 92


def decorate1790_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1790_0
def doubled1790_0(value: int) -> int:
    return value + value
def guarded1790_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1790_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 9:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1790_3(value: int) -> int:
    return value + 16


def decorate1790_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1790_3
def doubled1790_3(value: int) -> int:
    return value + value
def expand1790_4(values):
    for item in values:
        if item > 23:
            yield item * 17


def collect1790_4(values):
    return list(expand1790_4(values))
def traced1790_5(value: int) -> int:
    return value + 30


def decorate1790_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1790_5
def doubled1790_5(value: int) -> int:
    return value + value
def traced1790_6(value: int) -> int:
    return value + 37


def decorate1790_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1790_6
def doubled1790_6(value: int) -> int:
    return value + value
def traced1790_7(value: int) -> int:
    return value + 44


def decorate1790_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1790_7
def doubled1790_7(value: int) -> int:
    return value + value
