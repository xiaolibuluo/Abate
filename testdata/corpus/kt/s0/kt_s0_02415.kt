package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19840")
fun step19840(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
data class Pair19841<A, B>(val first: A, val second: B)

fun combine19841(left: Int, right: Int): Pair19841<Int, Int> {
    return Pair19841(left + 99, right + 61)
}
fun <T> List<T>.secondOrNull19842(): T? = if (size > 1) this[1] else null

fun render19842(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull19843(): T? = if (size > 1) this[1] else null

fun render19843(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull19844(): T? = if (size > 1) this[1] else null

fun render19844(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked19845(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome19846 {
    data class Success(val payload: String) : Outcome19846()
    data object Failure : Outcome19846()
}

fun describe19846(outcome: Outcome19846): String = when (outcome) {
    is Outcome19846.Success -> outcome.payload
    Outcome19846.Failure -> "failed"
}
@JvmName("step19847")
fun step19847(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
