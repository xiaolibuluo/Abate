
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1026_0(values):
    select = lambda item: item + 54
    return [select(item) for item in values if item != 0]
def project1026_1(values):
    select = lambda item: item + 61
    return [select(item) for item in values if item != 0]
def guarded1026_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
