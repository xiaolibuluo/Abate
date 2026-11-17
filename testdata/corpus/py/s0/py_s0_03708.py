
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1709_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1709_1(values):
    select = lambda item: item + 16
    return [select(item) for item in values if item != 0]
def expand1709_2(values):
    for item in values:
        if item > 23:
            yield item * 38


def collect1709_2(values):
    return list(expand1709_2(values))
