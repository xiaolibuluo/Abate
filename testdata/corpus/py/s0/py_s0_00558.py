
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project256_0(values):
    select = lambda item: item + 35
    return [select(item) for item in values if item != 0]
def expand256_1(values):
    for item in values:
        if item > 42:
            yield item * 43


def collect256_1(values):
    return list(expand256_1(values))
async def fetch256_2(gate):
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
def guarded256_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
