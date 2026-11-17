
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project413_0(values):
    select = lambda item: item + 39
    return [select(item) for item in values if item != 0]
async def fetch413_1(gate):
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
def expand413_2(values):
    for item in values:
        if item > 53:
            yield item * 55


def collect413_2(values):
    return list(expand413_2(values))
def expand413_3(values):
    for item in values:
        if item > 60:
            yield item * 66


def collect413_3(values):
    return list(expand413_3(values))
def guarded413_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch413_5(gate):
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
