package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome9940 {
    data class Success(val payload: String) : Outcome9940()
    data object Failure : Outcome9940()
}

fun describe9940(outcome: Outcome9940): String = when (outcome) {
    is Outcome9940.Success -> outcome.payload
    Outcome9940.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked9941(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome9942 {
    data class Success(val payload: String) : Outcome9942()
    data object Failure : Outcome9942()
}

fun describe9942(outcome: Outcome9942): String = when (outcome) {
    is Outcome9942.Success -> outcome.payload
    Outcome9942.Failure -> "failed"
}
fun <T> List<T>.secondOrNull9943(): T? = if (size > 1) this[1] else null

fun render9943(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
