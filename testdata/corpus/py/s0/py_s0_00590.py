
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project272_0(values):
    select = lambda item: item + 49
    return [select(item) for item in values if item != 0]
def guarded272_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project272_2(values):
    select = lambda item: item + 63
    return [select(item) for item in values if item != 0]
def summarize272_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 70:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded272_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced272_5(value: int) -> int:
    return value + 84


def decorate272_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate272_5
def doubled272_5(value: int) -> int:
    return value + value
def guarded272_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
