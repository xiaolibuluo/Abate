
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch532_0(gate):
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
def guarded532_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project532_2(values):
    select = lambda item: item + 48
    return [select(item) for item in values if item != 0]
async def fetch532_3(gate):
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
def traced532_4(value: int) -> int:
    return value + 62


def decorate532_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate532_4
def doubled532_4(value: int) -> int:
    return value + value
def guarded532_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch532_6(gate):
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
def expand532_7(values):
    for item in values:
        if item > 83:
            yield item * 28


def collect532_7(values):
    return list(expand532_7(values))
def project532_8(values):
    select = lambda item: item + 90
    return [select(item) for item in values if item != 0]
async def fetch532_9(gate):
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
