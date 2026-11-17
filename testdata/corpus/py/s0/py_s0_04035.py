
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1872_0(value: int) -> int:
    return value + 91


def decorate1872_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1872_0
def doubled1872_0(value: int) -> int:
    return value + value
def summarize1872_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 98:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1872_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 8:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1872_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1872_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
