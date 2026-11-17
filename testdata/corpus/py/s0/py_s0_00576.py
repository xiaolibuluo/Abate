
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand263_0(values):
    for item in values:
        if item > 29:
            yield item * 20


def collect263_0(values):
    return list(expand263_0(values))
def guarded263_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced263_2(value: int) -> int:
    return value + 43


def decorate263_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate263_2
def doubled263_2(value: int) -> int:
    return value + value
def project263_3(values):
    select = lambda item: item + 50
    return [select(item) for item in values if item != 0]
def summarize263_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 57:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
