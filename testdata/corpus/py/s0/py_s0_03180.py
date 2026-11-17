
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1471_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1471_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 26:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1471_2(values):
    for item in values:
        if item > 33:
            yield item * 19


def collect1471_2(values):
    return list(expand1471_2(values))
