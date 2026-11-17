package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4900<A, B>(val first: A, val second: B)

fun combine4900(left: Int, right: Int): Pair4900<Int, Int> {
    return Pair4900(left + 70, right + 23)
}
fun <T> List<T>.secondOrNull4901(): T? = if (size > 1) this[1] else null

fun render4901(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
