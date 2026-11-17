package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull1650(): T? = if (size > 1) this[1] else null

fun render1650(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull1651(): T? = if (size > 1) this[1] else null

fun render1651(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step1652")
fun step1652(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
data class Pair1653<A, B>(val first: A, val second: B)

fun combine1653(left: Int, right: Int): Pair1653<Int, Int> {
    return Pair1653(left + 37, right + 38)
}
fun <T> List<T>.secondOrNull1654(): T? = if (size > 1) this[1] else null

fun render1654(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
