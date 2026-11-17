package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13930 {
    data class Success(val payload: String) : Outcome13930()
    data object Failure : Outcome13930()
}

fun describe13930(outcome: Outcome13930): String = when (outcome) {
    is Outcome13930.Success -> outcome.payload
    Outcome13930.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13931(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked13932(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull13933(): T? = if (size > 1) this[1] else null

fun render13933(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13934 {
    data class Success(val payload: String) : Outcome13934()
    data object Failure : Outcome13934()
}

fun describe13934(outcome: Outcome13934): String = when (outcome) {
    is Outcome13934.Success -> outcome.payload
    Outcome13934.Failure -> "failed"
}
fun <T> List<T>.secondOrNull13935(): T? = if (size > 1) this[1] else null

fun render13935(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
