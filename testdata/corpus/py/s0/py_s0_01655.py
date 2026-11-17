
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced753_0(value: int) -> int:
    return value + 94


def decorate753_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate753_0
def doubled753_0(value: int) -> int:
    return value + value
def summarize753_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 101:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded753_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project753_3(values):
    select = lambda item: item + 18
    return [select(item) for item in values if item != 0]
def summarize753_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 25:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
