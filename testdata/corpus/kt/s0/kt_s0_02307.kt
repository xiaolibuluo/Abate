package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18940")
fun step18940(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
fun <T> List<T>.secondOrNull18941(): T? = if (size > 1) this[1] else null

fun render18941(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull18942(): T? = if (size > 1) this[1] else null

fun render18942(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome18943 {
    data class Success(val payload: String) : Outcome18943()
    data object Failure : Outcome18943()
}

fun describe18943(outcome: Outcome18943): String = when (outcome) {
    is Outcome18943.Success -> outcome.payload
    Outcome18943.Failure -> "failed"
}
data class Pair18944<A, B>(val first: A, val second: B)

fun combine18944(left: Int, right: Int): Pair18944<Int, Int> {
    return Pair18944(left + 17, right + 13)
}
fun <T> List<T>.secondOrNull18945(): T? = if (size > 1) this[1] else null

fun render18945(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked18946(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull18947(): T? = if (size > 1) this[1] else null

fun render18947(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull18948(): T? = if (size > 1) this[1] else null

fun render18948(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
