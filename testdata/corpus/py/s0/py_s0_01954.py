
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand892_0(values):
    for item in values:
        if item > 58:
            yield item * 31


def collect892_0(values):
    return list(expand892_0(values))
async def fetch892_1(gate):
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
def expand892_2(values):
    for item in values:
        if item > 72:
            yield item * 53


def collect892_2(values):
    return list(expand892_2(values))
def guarded892_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
