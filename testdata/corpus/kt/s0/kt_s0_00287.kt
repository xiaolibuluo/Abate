package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step2380")
fun step2380(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
data class Pair2381<A, B>(val first: A, val second: B)

fun combine2381(left: Int, right: Int): Pair2381<Int, Int> {
    return Pair2381(left + 99, right + 39)
}
@JvmName("step2382")
fun step2382(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
fun <T> List<T>.secondOrNull2383(): T? = if (size > 1) this[1] else null

fun render2383(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked2384(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair2385<A, B>(val first: A, val second: B)

fun combine2385(left: Int, right: Int): Pair2385<Int, Int> {
    return Pair2385(left + 30, right + 22)
}
fun <T, R> List<T>.mapChecked2386(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step2387")
fun step2387(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
data class Pair2388<A, B>(val first: A, val second: B)

fun combine2388(left: Int, right: Int): Pair2388<Int, Int> {
    return Pair2388(left + 51, right + 55)
}
@JvmName("step2389")
fun step2389(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
object Registry23810 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23810 = "registry-65"
    }
}
