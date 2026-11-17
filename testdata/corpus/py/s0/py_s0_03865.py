
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1785_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1785_1(values):
    for item in values:
        if item > 34:
            yield item * 10


def collect1785_1(values):
    return list(expand1785_1(values))
def traced1785_2(value: int) -> int:
    return value + 41


def decorate1785_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1785_2
def doubled1785_2(value: int) -> int:
    return value + value
def guarded1785_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1785_4(values):
    select = lambda item: item + 55
    return [select(item) for item in values if item != 0]
def summarize1785_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1785_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1785_7(value: int) -> int:
    return value + 76


def decorate1785_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1785_7
def doubled1785_7(value: int) -> int:
    return value + value
