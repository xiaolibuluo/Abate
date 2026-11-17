
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1289_0(gate):
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
def expand1289_1(values):
    for item in values:
        if item > 85:
            yield item * 15


def collect1289_1(values):
    return list(expand1289_1(values))
def project1289_2(values):
    select = lambda item: item + 92
    return [select(item) for item in values if item != 0]
def expand1289_3(values):
    for item in values:
        if item > 99:
            yield item * 37


def collect1289_3(values):
    return list(expand1289_3(values))
def guarded1289_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1289_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
