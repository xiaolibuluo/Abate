
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1326_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1326_1(values):
    select = lambda item: item + 81
    return [select(item) for item in values if item != 0]
def expand1326_2(values):
    for item in values:
        if item > 88:
            yield item * 41


def collect1326_2(values):
    return list(expand1326_2(values))
def expand1326_3(values):
    for item in values:
        if item > 95:
            yield item * 52


def collect1326_3(values):
    return list(expand1326_3(values))
