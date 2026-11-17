package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair17120<A, B>(val first: A, val second: B)

fun combine17120(left: Int, right: Int): Pair17120<Int, Int> {
    return Pair17120(left + 48, right + 37)
}
fun <T> List<T>.secondOrNull17121(): T? = if (size > 1) this[1] else null

fun render17121(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair17122<A, B>(val first: A, val second: B)

fun combine17122(left: Int, right: Int): Pair17122<Int, Int> {
    return Pair17122(left + 62, right + 59)
}
