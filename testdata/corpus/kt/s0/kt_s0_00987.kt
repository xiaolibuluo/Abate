package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8090 {
    data class Success(val payload: String) : Outcome8090()
    data object Failure : Outcome8090()
}

fun describe8090(outcome: Outcome8090): String = when (outcome) {
    is Outcome8090.Success -> outcome.payload
    Outcome8090.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked8091(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8092 {
    data class Success(val payload: String) : Outcome8092()
    data object Failure : Outcome8092()
}

fun describe8092(outcome: Outcome8092): String = when (outcome) {
    is Outcome8092.Success -> outcome.payload
    Outcome8092.Failure -> "failed"
}
fun <T> List<T>.secondOrNull8093(): T? = if (size > 1) this[1] else null

fun render8093(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair8094<A, B>(val first: A, val second: B)

fun combine8094(left: Int, right: Int): Pair8094<Int, Int> {
    return Pair8094(left + 74, right + 43)
}
fun <T> List<T>.secondOrNull8095(): T? = if (size > 1) this[1] else null

fun render8095(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
