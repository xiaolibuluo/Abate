package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21230(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull21231(): T? = if (size > 1) this[1] else null

fun render21231(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21232<A, B>(val first: A, val second: B)

fun combine21232(left: Int, right: Int): Pair21232<Int, Int> {
    return Pair21232(left + 70, right + 69)
}
data class Pair21233<A, B>(val first: A, val second: B)

fun combine21233(left: Int, right: Int): Pair21233<Int, Int> {
    return Pair21233(left + 77, right + 19)
}
object Registry21234 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21234 = "registry-84"
    }
}
@JvmName("step21235")
fun step21235(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
data class Pair21236<A, B>(val first: A, val second: B)

fun combine21236(left: Int, right: Int): Pair21236<Int, Int> {
    return Pair21236(left + 98, right + 52)
}
object Registry21237 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21237 = "registry-8"
    }
}
