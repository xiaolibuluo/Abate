
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2552_0(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
def expand2552_1(values):
    for item in values:
        if item > 14:
            yield item * 11


def collect2552_1(values):
    return list(expand2552_1(values))
def expand2552_2(values):
    for item in values:
        if item > 21:
            yield item * 22


def collect2552_2(values):
    return list(expand2552_2(values))
async def fetch2552_3(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def guarded2552_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2552_5(values):
    for item in values:
        if item > 42:
            yield item * 55


def collect2552_5(values):
    return list(expand2552_5(values))
