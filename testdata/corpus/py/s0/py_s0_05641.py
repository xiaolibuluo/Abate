
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2597_0(value: int) -> int:
    return value + 10


def decorate2597_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2597_0
def doubled2597_0(value: int) -> int:
    return value + value
def summarize2597_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 17:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2597_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2597_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
