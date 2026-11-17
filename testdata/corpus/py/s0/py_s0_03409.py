
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1576_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 26:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1576_1(values):
    for item in values:
        if item > 33:
            yield item * 11


def collect1576_1(values):
    return list(expand1576_1(values))
def traced1576_2(value: int) -> int:
    return value + 40


def decorate1576_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1576_2
def doubled1576_2(value: int) -> int:
    return value + value
def guarded1576_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1576_4(value: int) -> int:
    return value + 54


def decorate1576_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1576_4
def doubled1576_4(value: int) -> int:
    return value + value
