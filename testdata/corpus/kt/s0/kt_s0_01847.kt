package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15180<A, B>(val first: A, val second: B)

fun combine15180(left: Int, right: Int): Pair15180<Int, Int> {
    return Pair15180(left + 48, right + 21)
}
data class Pair15181<A, B>(val first: A, val second: B)

fun combine15181(left: Int, right: Int): Pair15181<Int, Int> {
    return Pair15181(left + 55, right + 32)
}
fun <T> List<T>.secondOrNull15182(): T? = if (size > 1) this[1] else null

fun render15182(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
