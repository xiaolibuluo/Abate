
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2258_0(value: int) -> int:
    return value + 65


def decorate2258_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2258_0
def doubled2258_0(value: int) -> int:
    return value + value
def summarize2258_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2258_2(values):
    for item in values:
        if item > 79:
            yield item * 38


def collect2258_2(values):
    return list(expand2258_2(values))
def guarded2258_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2258_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
