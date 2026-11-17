package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull16620(): T? = if (size > 1) this[1] else null

fun render16620(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16621<A, B>(val first: A, val second: B)

fun combine16621(left: Int, right: Int): Pair16621<Int, Int> {
    return Pair16621(left + 84, right + 64)
}
fun <T> List<T>.secondOrNull16622(): T? = if (size > 1) this[1] else null

fun render16622(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
