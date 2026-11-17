
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2118_0(value: int) -> int:
    return value + 88


def decorate2118_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2118_0
def doubled2118_0(value: int) -> int:
    return value + value
def guarded2118_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2118_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 5}
def summarize2118_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 12:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2118_4(value: int) -> int:
    return value + 19


def decorate2118_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2118_4
def doubled2118_4(value: int) -> int:
    return value + value
def guarded2118_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
