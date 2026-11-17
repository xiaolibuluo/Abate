package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull3530(): T? = if (size > 1) this[1] else null

fun render3530(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair3531<A, B>(val first: A, val second: B)

fun combine3531(left: Int, right: Int): Pair3531<Int, Int> {
    return Pair3531(left + 42, right + 51)
}
fun <T> List<T>.secondOrNull3532(): T? = if (size > 1) this[1] else null

fun render3532(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
