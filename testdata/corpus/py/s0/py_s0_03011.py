
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1381_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1381_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1381_2(value: int) -> int:
    return value + 27


def decorate1381_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1381_2
def doubled1381_2(value: int) -> int:
    return value + value
def summarize1381_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 34:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
