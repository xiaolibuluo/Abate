
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1958_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1958_1(value: int) -> int:
    return value + 52


def decorate1958_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1958_1
def doubled1958_1(value: int) -> int:
    return value + value
def project1958_2(values):
    select = lambda item: item + 59
    return [select(item) for item in values if item != 0]
def guarded1958_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
