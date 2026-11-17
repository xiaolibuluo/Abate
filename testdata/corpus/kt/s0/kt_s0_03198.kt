package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair26290<A, B>(val first: A, val second: B)

fun combine26290(left: Int, right: Int): Pair26290<Int, Int> {
    return Pair26290(left + 38, right + 51)
}
fun <T> List<T>.secondOrNull26291(): T? = if (size > 1) this[1] else null

fun render26291(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
