package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull19500(): T? = if (size > 1) this[1] else null

fun render19500(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair19501<A, B>(val first: A, val second: B)

fun combine19501(left: Int, right: Int): Pair19501<Int, Int> {
    return Pair19501(left + 45, right + 67)
}
