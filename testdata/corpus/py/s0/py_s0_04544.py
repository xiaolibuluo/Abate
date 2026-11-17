
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2112_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2112_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2112_2(values):
    for item in values:
        if item > 24:
            yield item * 53


def collect2112_2(values):
    return list(expand2112_2(values))
def traced2112_3(value: int) -> int:
    return value + 31


def decorate2112_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2112_3
def doubled2112_3(value: int) -> int:
    return value + value
def project2112_4(values):
    select = lambda item: item + 38
    return [select(item) for item in values if item != 0]
