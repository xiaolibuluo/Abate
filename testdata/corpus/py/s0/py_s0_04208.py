
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1962_0(values):
    for item in values:
        if item > 97:
            yield item * 18


def collect1962_0(values):
    return list(expand1962_0(values))
async def fetch1962_1(gate):
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
def project1962_2(values):
    select = lambda item: item + 14
    return [select(item) for item in values if item != 0]
def guarded1962_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1962_4(values):
    for item in values:
        if item > 28:
            yield item * 62


def collect1962_4(values):
    return list(expand1962_4(values))
