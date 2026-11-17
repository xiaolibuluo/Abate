package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18320<A, B>(val first: A, val second: B)

fun combine18320(left: Int, right: Int): Pair18320<Int, Int> {
    return Pair18320(left + 56, right + 23)
}
object Registry18321 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18321 = "registry-63"
    }
}
fun <T, R> List<T>.mapChecked18322(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked18323(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked18324(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull18325(): T? = if (size > 1) this[1] else null

fun render18325(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step18326")
fun step18326(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 98
}
