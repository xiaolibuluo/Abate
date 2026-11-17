package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18420 {
    data class Success(val payload: String) : Outcome18420()
    data object Failure : Outcome18420()
}

fun describe18420(outcome: Outcome18420): String = when (outcome) {
    is Outcome18420.Success -> outcome.payload
    Outcome18420.Failure -> "failed"
}
object Registry18421 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18421 = "registry-96"
    }
}
fun <T> List<T>.secondOrNull18422(): T? = if (size > 1) this[1] else null

fun render18422(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18423<A, B>(val first: A, val second: B)

fun combine18423(left: Int, right: Int): Pair18423<Int, Int> {
    return Pair18423(left + 13, right + 65)
}
fun <T, R> List<T>.mapChecked18424(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked18425(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18426 {
    data class Success(val payload: String) : Outcome18426()
    data object Failure : Outcome18426()
}

fun describe18426(outcome: Outcome18426): String = when (outcome) {
    is Outcome18426.Success -> outcome.payload
    Outcome18426.Failure -> "failed"
}
@JvmName("step18427")
fun step18427(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
