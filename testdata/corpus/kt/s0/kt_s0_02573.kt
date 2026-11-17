package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull21200(): T? = if (size > 1) this[1] else null

fun render21200(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome21201 {
    data class Success(val payload: String) : Outcome21201()
    data object Failure : Outcome21201()
}

fun describe21201(outcome: Outcome21201): String = when (outcome) {
    is Outcome21201.Success -> outcome.payload
    Outcome21201.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21202(): T? = if (size > 1) this[1] else null

fun render21202(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull21203(): T? = if (size > 1) this[1] else null

fun render21203(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked21204(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair21205<A, B>(val first: A, val second: B)

fun combine21205(left: Int, right: Int): Pair21205<Int, Int> {
    return Pair21205(left + 52, right + 20)
}
