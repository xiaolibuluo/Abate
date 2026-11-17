
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch841_0(gate):
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
def expand841_1(values):
    for item in values:
        if item > 81:
            yield item * 51


def collect841_1(values):
    return list(expand841_1(values))
def guarded841_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project841_3(values):
    select = lambda item: item + 95
    return [select(item) for item in values if item != 0]
