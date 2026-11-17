
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2213_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2213_1(values):
    for item in values:
        if item > 69:
            yield item * 17


def collect2213_1(values):
    return list(expand2213_1(values))
def summarize2213_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 76:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2213_3(value: int) -> int:
    return value + 83


def decorate2213_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2213_3
def doubled2213_3(value: int) -> int:
    return value + value
def traced2213_4(value: int) -> int:
    return value + 90


def decorate2213_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2213_4
def doubled2213_4(value: int) -> int:
    return value + value
