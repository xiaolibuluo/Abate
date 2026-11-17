package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull13430(): T? = if (size > 1) this[1] else null

fun render13430(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair13431<A, B>(val first: A, val second: B)

fun combine13431(left: Int, right: Int): Pair13431<Int, Int> {
    return Pair13431(left + 11, right + 27)
}
sealed class Outcome13432 {
    data class Success(val payload: String) : Outcome13432()
    data object Failure : Outcome13432()
}

fun describe13432(outcome: Outcome13432): String = when (outcome) {
    is Outcome13432.Success -> outcome.payload
    Outcome13432.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13433(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome13434 {
    data class Success(val payload: String) : Outcome13434()
    data object Failure : Outcome13434()
}

fun describe13434(outcome: Outcome13434): String = when (outcome) {
    is Outcome13434.Success -> outcome.payload
    Outcome13434.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13435(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
