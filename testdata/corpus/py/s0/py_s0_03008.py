
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1378_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1378_1(value: int) -> int:
    return value + 78


def decorate1378_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1378_1
def doubled1378_1(value: int) -> int:
    return value + value
def guarded1378_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1378_3(value: int) -> int:
    return value + 92


def decorate1378_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1378_3
def doubled1378_3(value: int) -> int:
    return value + value
