package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15950<A, B>(val first: A, val second: B)

fun combine15950(left: Int, right: Int): Pair15950<Int, Int> {
    return Pair15950(left + 79, right + 11)
}
fun <T> List<T>.secondOrNull15951(): T? = if (size > 1) this[1] else null

fun render15951(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15952<A, B>(val first: A, val second: B)

fun combine15952(left: Int, right: Int): Pair15952<Int, Int> {
    return Pair15952(left + 93, right + 33)
}
fun <T, R> List<T>.mapChecked15953(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step15954")
fun step15954(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
fun <T, R> List<T>.mapChecked15955(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked15956(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry15957 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15957 = "registry-31"
    }
}
data class Pair15958<A, B>(val first: A, val second: B)

fun combine15958(left: Int, right: Int): Pair15958<Int, Int> {
    return Pair15958(left + 38, right + 38)
}
