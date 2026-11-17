package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8050<A, B>(val first: A, val second: B)

fun combine8050(left: Int, right: Int): Pair8050<Int, Int> {
    return Pair8050(left + 91, right + 32)
}
fun <T, R> List<T>.mapChecked8051(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8052<A, B>(val first: A, val second: B)

fun combine8052(left: Int, right: Int): Pair8052<Int, Int> {
    return Pair8052(left + 8, right + 54)
}
fun <T> List<T>.secondOrNull8053(): T? = if (size > 1) this[1] else null

fun render8053(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair8054<A, B>(val first: A, val second: B)

fun combine8054(left: Int, right: Int): Pair8054<Int, Int> {
    return Pair8054(left + 22, right + 15)
}
@JvmName("step8055")
fun step8055(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 29
}
@JvmName("step8056")
fun step8056(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
