
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded57_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced57_1(value: int) -> int:
    return value + 74


def decorate57_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate57_1
def doubled57_1(value: int) -> int:
    return value + value
def summarize57_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 81:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
