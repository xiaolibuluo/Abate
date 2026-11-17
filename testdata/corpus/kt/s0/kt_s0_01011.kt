package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8290<A, B>(val first: A, val second: B)

fun combine8290(left: Int, right: Int): Pair8290<Int, Int> {
    return Pair8290(left + 15, right + 17)
}
fun <T> List<T>.secondOrNull8291(): T? = if (size > 1) this[1] else null

fun render8291(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
