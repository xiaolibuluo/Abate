
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2058_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2058_1(value: int) -> int:
    return value + 91


def decorate2058_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2058_1
def doubled2058_1(value: int) -> int:
    return value + value
def guarded2058_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2058_3(values):
    select = lambda item: item + 8
    return [select(item) for item in values if item != 0]
