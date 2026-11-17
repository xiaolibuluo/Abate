
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2598_0(values):
    select = lambda item: item + 23
    return [select(item) for item in values if item != 0]
def guarded2598_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2598_2(value: int) -> int:
    return value + 37


def decorate2598_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2598_2
def doubled2598_2(value: int) -> int:
    return value + value
def summarize2598_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 44:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2598_4(value: int) -> int:
    return value + 51


def decorate2598_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2598_4
def doubled2598_4(value: int) -> int:
    return value + value
