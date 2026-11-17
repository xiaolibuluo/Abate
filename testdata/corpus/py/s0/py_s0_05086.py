
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2365_0(value: int) -> int:
    return value + 98


def decorate2365_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2365_0
def doubled2365_0(value: int) -> int:
    return value + value
def guarded2365_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2365_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2365_3(value: int) -> int:
    return value + 22


def decorate2365_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2365_3
def doubled2365_3(value: int) -> int:
    return value + value
def project2365_4(values):
    select = lambda item: item + 29
    return [select(item) for item in values if item != 0]
def guarded2365_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2365_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 43:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
