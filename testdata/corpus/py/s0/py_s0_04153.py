
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1932_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 95:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1932_1(value: int) -> int:
    return value + 5


def decorate1932_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1932_1
def doubled1932_1(value: int) -> int:
    return value + value
def summarize1932_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 12:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1932_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1932_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 26:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1932_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1932_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
