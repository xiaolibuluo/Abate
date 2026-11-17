package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry18380 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18380 = "registry-37"
    }
}
fun <T> List<T>.secondOrNull18381(): T? = if (size > 1) this[1] else null

fun render18381(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome18382 {
    data class Success(val payload: String) : Outcome18382()
    data object Failure : Outcome18382()
}

fun describe18382(outcome: Outcome18382): String = when (outcome) {
    is Outcome18382.Success -> outcome.payload
    Outcome18382.Failure -> "failed"
}
fun <T> List<T>.secondOrNull18383(): T? = if (size > 1) this[1] else null

fun render18383(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked18384(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry18385 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18385 = "registry-72"
    }
}
fun <T> List<T>.secondOrNull18386(): T? = if (size > 1) this[1] else null

fun render18386(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18387<A, B>(val first: A, val second: B)

fun combine18387(left: Int, right: Int): Pair18387<Int, Int> {
    return Pair18387(left + 86, right + 20)
}
