package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23520<A, B>(val first: A, val second: B)

fun combine23520(left: Int, right: Int): Pair23520<Int, Int> {
    return Pair23520(left + 26, right + 64)
}
@JvmName("step23521")
fun step23521(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
fun <T> List<T>.secondOrNull23522(): T? = if (size > 1) this[1] else null

fun render23522(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull23523(): T? = if (size > 1) this[1] else null

fun render23523(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
