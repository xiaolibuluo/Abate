package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull9870(): T? = if (size > 1) this[1] else null

fun render9870(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked9871(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome9872 {
    data class Success(val payload: String) : Outcome9872()
    data object Failure : Outcome9872()
}

fun describe9872(outcome: Outcome9872): String = when (outcome) {
    is Outcome9872.Success -> outcome.payload
    Outcome9872.Failure -> "failed"
}
