package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair2690<A, B>(val first: A, val second: B)

fun combine2690(left: Int, right: Int): Pair2690<Int, Int> {
    return Pair2690(left + 10, right + 62)
}
object Registry2691 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2691 = "registry-17"
    }
}
fun <T> List<T>.secondOrNull2692(): T? = if (size > 1) this[1] else null

fun render2692(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step2693")
fun step2693(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
data class Pair2694<A, B>(val first: A, val second: B)

fun combine2694(left: Int, right: Int): Pair2694<Int, Int> {
    return Pair2694(left + 38, right + 45)
}
fun <T, R> List<T>.mapChecked2695(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
