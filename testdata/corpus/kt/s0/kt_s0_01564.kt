package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked12940(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12941 {
    data class Success(val payload: String) : Outcome12941()
    data object Failure : Outcome12941()
}

fun describe12941(outcome: Outcome12941): String = when (outcome) {
    is Outcome12941.Success -> outcome.payload
    Outcome12941.Failure -> "failed"
}
fun <T> List<T>.secondOrNull12942(): T? = if (size > 1) this[1] else null

fun render12942(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair12943<A, B>(val first: A, val second: B)

fun combine12943(left: Int, right: Int): Pair12943<Int, Int> {
    return Pair12943(left + 67, right + 11)
}
data class Pair12944<A, B>(val first: A, val second: B)

fun combine12944(left: Int, right: Int): Pair12944<Int, Int> {
    return Pair12944(left + 74, right + 22)
}
sealed class Outcome12945 {
    data class Success(val payload: String) : Outcome12945()
    data object Failure : Outcome12945()
}

fun describe12945(outcome: Outcome12945): String = when (outcome) {
    is Outcome12945.Success -> outcome.payload
    Outcome12945.Failure -> "failed"
}
