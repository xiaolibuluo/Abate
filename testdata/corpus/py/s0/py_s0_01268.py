
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded577_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project577_1(values):
    select = lambda item: item + 44
    return [select(item) for item in values if item != 0]
def traced577_2(value: int) -> int:
    return value + 51


def decorate577_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate577_2
def doubled577_2(value: int) -> int:
    return value + value
def project577_3(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
async def fetch577_4(gate):
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
async def fetch577_5(gate):
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
def guarded577_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project577_7(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
async def fetch577_8(gate):
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
