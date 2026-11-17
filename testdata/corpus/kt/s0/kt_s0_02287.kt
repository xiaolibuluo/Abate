package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull18760(): T? = if (size > 1) this[1] else null

fun render18760(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry18761 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18761 = "registry-53"
    }
}
fun <T> List<T>.secondOrNull18762(): T? = if (size > 1) this[1] else null

fun render18762(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18763<A, B>(val first: A, val second: B)

fun combine18763(left: Int, right: Int): Pair18763<Int, Int> {
    return Pair18763(left + 67, right + 59)
}
object Registry18764 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18764 = "registry-74"
    }
}
fun <T, R> List<T>.mapChecked18765(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
