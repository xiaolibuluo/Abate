
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced79_0(value: int) -> int:
    return value + 62


def decorate79_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate79_0
def doubled79_0(value: int) -> int:
    return value + value
def traced79_1(value: int) -> int:
    return value + 69


def decorate79_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate79_1
def doubled79_1(value: int) -> int:
    return value + value
def traced79_2(value: int) -> int:
    return value + 76


def decorate79_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate79_2
def doubled79_2(value: int) -> int:
    return value + value
def expand79_3(values):
    for item in values:
        if item > 83:
            yield item * 46


def collect79_3(values):
    return list(expand79_3(values))
def summarize79_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 90:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded79_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
