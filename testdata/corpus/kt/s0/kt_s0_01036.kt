package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step8520")
fun step8520(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
data class Pair8521<A, B>(val first: A, val second: B)

fun combine8521(left: Int, right: Int): Pair8521<Int, Int> {
    return Pair8521(left + 30, right + 67)
}
fun <T> List<T>.secondOrNull8522(): T? = if (size > 1) this[1] else null

fun render8522(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair8523<A, B>(val first: A, val second: B)

fun combine8523(left: Int, right: Int): Pair8523<Int, Int> {
    return Pair8523(left + 44, right + 28)
}
@JvmName("step8524")
fun step8524(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
@JvmName("step8525")
fun step8525(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
fun <T> List<T>.secondOrNull8526(): T? = if (size > 1) this[1] else null

fun render8526(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
