package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull9050(): T? = if (size > 1) this[1] else null

fun render9050(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step9051")
fun step9051(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
fun <T, R> List<T>.mapChecked9052(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull9053(): T? = if (size > 1) this[1] else null

fun render9053(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull9054(): T? = if (size > 1) this[1] else null

fun render9054(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome9055 {
    data class Success(val payload: String) : Outcome9055()
    data object Failure : Outcome9055()
}

fun describe9055(outcome: Outcome9055): String = when (outcome) {
    is Outcome9055.Success -> outcome.payload
    Outcome9055.Failure -> "failed"
}
