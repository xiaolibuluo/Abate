package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull21690(): T? = if (size > 1) this[1] else null

fun render21690(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step21691")
fun step21691(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
data class Pair21692<A, B>(val first: A, val second: B)

fun combine21692(left: Int, right: Int): Pair21692<Int, Int> {
    return Pair21692(left + 86, right + 25)
}
data class Pair21693<A, B>(val first: A, val second: B)

fun combine21693(left: Int, right: Int): Pair21693<Int, Int> {
    return Pair21693(left + 93, right + 36)
}
@JvmName("step21694")
fun step21694(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
@JvmName("step21695")
fun step21695(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
fun <T, R> List<T>.mapChecked21696(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
