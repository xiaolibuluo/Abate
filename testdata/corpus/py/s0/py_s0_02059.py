
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand941_0(values):
    for item in values:
        if item > 16:
            yield item * 69


def collect941_0(values):
    return list(expand941_0(values))
async def fetch941_1(gate):
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
def expand941_2(values):
    for item in values:
        if item > 30:
            yield item * 30


def collect941_2(values):
    return list(expand941_2(values))
def guarded941_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded941_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
