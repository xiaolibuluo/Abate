
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2366_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2366_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
def traced2366_2(value: int) -> int:
    return value + 28


def decorate2366_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2366_2
def doubled2366_2(value: int) -> int:
    return value + value
def summarize2366_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 35:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2366_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 42:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2366_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
