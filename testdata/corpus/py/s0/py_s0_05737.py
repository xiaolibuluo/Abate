
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2640_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 84:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2640_1(value: int) -> int:
    return value + 91


def decorate2640_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2640_1
def doubled2640_1(value: int) -> int:
    return value + value
def guarded2640_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
