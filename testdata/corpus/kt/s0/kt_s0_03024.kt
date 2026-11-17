package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull24850(): T? = if (size > 1) this[1] else null

fun render24850(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24851<A, B>(val first: A, val second: B)

fun combine24851(left: Int, right: Int): Pair24851<Int, Int> {
    return Pair24851(left + 16, right + 30)
}
