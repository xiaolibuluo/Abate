
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1428_0(values):
    for item in values:
        if item > 42:
            yield item * 62


def collect1428_0(values):
    return list(expand1428_0(values))
def guarded1428_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1428_2(values):
    for item in values:
        if item > 56:
            yield item * 23


def collect1428_2(values):
    return list(expand1428_2(values))
async def fetch1428_3(gate):
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
async def fetch1428_4(gate):
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
