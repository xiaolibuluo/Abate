
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced124_0(value: int) -> int:
    return value + 65


def decorate124_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate124_0
def doubled124_0(value: int) -> int:
    return value + value
def project124_1(values):
    select = lambda item: item + 72
    return [select(item) for item in values if item != 0]
def traced124_2(value: int) -> int:
    return value + 79


def decorate124_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate124_2
def doubled124_2(value: int) -> int:
    return value + value
def guarded124_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced124_4(value: int) -> int:
    return value + 93


def decorate124_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate124_4
def doubled124_4(value: int) -> int:
    return value + value
def summarize124_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
