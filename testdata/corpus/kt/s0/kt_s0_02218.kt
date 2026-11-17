package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull18260(): T? = if (size > 1) this[1] else null

fun render18260(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18261<A, B>(val first: A, val second: B)

fun combine18261(left: Int, right: Int): Pair18261<Int, Int> {
    return Pair18261(left + 82, right + 53)
}
fun <T, R> List<T>.mapChecked18262(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked18263(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step18264")
fun step18264(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
sealed class Outcome18265 {
    data class Success(val payload: String) : Outcome18265()
    data object Failure : Outcome18265()
}

fun describe18265(outcome: Outcome18265): String = when (outcome) {
    is Outcome18265.Success -> outcome.payload
    Outcome18265.Failure -> "failed"
}
