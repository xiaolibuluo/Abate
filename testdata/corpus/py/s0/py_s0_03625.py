
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1669_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1669_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1669_2(values):
    select = lambda item: item + 85
    return [select(item) for item in values if item != 0]
def expand1669_3(values):
    for item in values:
        if item > 92:
            yield item * 13


def collect1669_3(values):
    return list(expand1669_3(values))
