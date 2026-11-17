package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked18430(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18431 {
    data class Success(val payload: String) : Outcome18431()
    data object Failure : Outcome18431()
}

fun describe18431(outcome: Outcome18431): String = when (outcome) {
    is Outcome18431.Success -> outcome.payload
    Outcome18431.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked18432(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair18433<A, B>(val first: A, val second: B)

fun combine18433(left: Int, right: Int): Pair18433<Int, Int> {
    return Pair18433(left + 26, right + 11)
}
fun <T> List<T>.secondOrNull18434(): T? = if (size > 1) this[1] else null

fun render18434(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome18435 {
    data class Success(val payload: String) : Outcome18435()
    data object Failure : Outcome18435()
}

fun describe18435(outcome: Outcome18435): String = when (outcome) {
    is Outcome18435.Success -> outcome.payload
    Outcome18435.Failure -> "failed"
}
@JvmName("step18436")
fun step18436(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
object Registry18437 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18437 = "registry-54"
    }
}
data class Pair18438<A, B>(val first: A, val second: B)

fun combine18438(left: Int, right: Int): Pair18438<Int, Int> {
    return Pair18438(left + 61, right + 66)
}
